#!/usr/bin/env bash


cat > ${domain}.conf <<eof
server {
  listen              80;
	server_name       ${domain};
	rewrite ^(.*)\$ https://\$host\$1 permanent;
}


server {
	  listen       443 ssl;
    server_name  ${domain};
    ssl_certificate  /etc/letsencrypt/live/${domain}/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/${domain}/privkey.pem;
    ssl_session_cache    shared:SSL:1m;
    ssl_session_timeout  5m;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4:!DH:!DHE;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_prefer_server_ciphers  on;


    location = / {
    root  /;
	  proxy_pass  http://my.mhmarkets.au;
		proxy_set_header Host ${domain};
 	  proxy_set_header X-Forwarded-Host  ${domain};
    proxy_set_header X-Forwarded-Server ${domain};
	  proxy_set_header  X-Real-IP  \$remote_addr;
    proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto  \$scheme;
	  #限制每ip每秒不超过20个请求，漏桶数burst为5
    #brust的意思就是，如果第1秒、2,3,4秒请求为19个，
    #第5秒的请求为25个是被允许的。
    #但是如果你第1秒就25个请求，第2秒超过20的请求返回503错误。
    #nodelay，如果不设置该选项，严格使用平均速率限制请求数，
    #第1秒25个请求时，5个请求放到第2秒执行，
    #设置nodelay，25个请求将在第1秒执行。
	  limit_req zone=static nodelay;
		proxy_set_header remote-user-ip \$remote_addr;
		#set_real_ip_from 0.0.0.0/0;
		real_ip_header X-Forwarded-For;
	  if (\$http_user_agent ~* "(mobile|nokia|iphone|ipad|android|samsung|htc|blackberry)") {
                rewrite ^/(.*)\$ https://mobile.mhmarkets.au\$request_uri redirect;
    }
  }
	#JS和CSS缓存时间设置,图片缓存时间设置
  location ~ \.(gif|jpg|bmp|js|css|woff|woff2|ttf|json|swf|ico|png)\$ {
  expires 30d;
  valid_referers  none blocked ${domain};
  if (\$invalid_referer) {
    #rewrite ^/ http://\$host/logo.png;
    return 403;
  }
	proxy_pass  http://my.mhmarkets.au;
	#limit_req zone=static burst=3 nodelay;
	proxy_set_header Host ${domain};
 	proxy_set_header X-Forwarded-Host ${domain};
  proxy_set_header X-Forwarded-Server ${domain};
  proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
  proxy_set_header  X-Real-IP  \$remote_addr;
	proxy_set_header X-Forwarded-Proto  \$scheme;
  proxy_set_header remote-user-ip \$remote_addr;
  }

	location /login/  {
  proxy_pass  http://my.mhmarkets.au;
  limit_req zone=static burst=200 nodelay;
  client_max_body_size  50m;
	proxy_connect_timeout 900;
  proxy_send_timeout 900;
  proxy_read_timeout 900;
	send_timeout 900;
	proxy_set_header Host ${domain};
  proxy_set_header X-Forwarded-Host ${domain};
  proxy_set_header X-Forwarded-Server ${domain};
  proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
  proxy_set_header  X-Real-IP  \$remote_addr;
	proxy_set_header X-Forwarded-Proto  \$scheme;
	proxy_set_header remote-user-ip \$remote_addr;
	#set_real_ip_from 0.0.0.0/0;
	real_ip_header X-Forwarded-For;
	if (\$http_user_agent ~* "(mobile|nokia|iphone|ipad|android|samsung|htc|blackberry)") {
    rewrite ^/(.*)\$ https://mobile.mhmarkets.au\$request_uri redirect;
	  }
  }

	location /register/  {
	proxy_pass  http://my.mhmarkets.au;
	limit_req zone=static burst=200 nodelay;
	client_max_body_size  50m;
	proxy_connect_timeout 900;
  proxy_send_timeout 900;
  proxy_read_timeout 900;
	send_timeout 900;
	proxy_set_header Host ${domain};
  proxy_set_header X-Forwarded-Host ${domain};
  proxy_set_header X-Forwarded-Server ${domain};
  proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
  proxy_set_header  X-Real-IP  \$remote_addr;
	proxy_set_header X-Forwarded-Proto \$scheme;
	proxy_set_header remote-user-ip \$remote_addr;
	#set_real_ip_from 0.0.0.0/0;
	real_ip_header X-Forwarded-For;
	if (\$http_user_agent ~* "(mobile|nokia|iphone|ipad|android|samsung|htc|blackberry)") {
             rewrite ^/(.*)\$ https://mobile.mhmarkets.au/\$1 permanent;
	  }
  }


  location /regActive/  {
	proxy_pass  http://my.mhmarkets.au;
	limit_req zone=static burst=200 nodelay;
	client_max_body_size  50m;
	proxy_connect_timeout 900;
  proxy_send_timeout 900;
  proxy_read_timeout 900;
	send_timeout 900;
	proxy_set_header Host ${domain};
  proxy_set_header X-Forwarded-Host ${domain};
  proxy_set_header X-Forwarded-Server ${domain};
  proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
  proxy_set_header  X-Real-IP  \$remote_addr;
	proxy_set_header X-Forwarded-Proto  \$scheme;
	proxy_set_header remote-user-ip \$remote_addr;
	#set_real_ip_from 0.0.0.0/0;
	real_ip_header X-Forwarded-For;
	if (\$http_user_agent ~* "(mobile|nokia|iphone|ipad|android|samsung|htc|blackberry)") {
                rewrite ^/(.*)\$ https://mobile.mhmarkets.au/\$1 permanent;
			}
  }

	location /api/ {
            proxy_pass http://8.222.203.254:8080/api/;
  }

	location / {
	  try_files \$uri \$uri/ @rewrites;
	  index index.html;
	}

	location @rewrites {
		  rewrite ^.*\$ / last;
	  }
}
eof