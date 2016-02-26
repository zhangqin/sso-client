<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="Cache-Control" content="max-age=3" />
  </head>
<body>
{{ if .ticket  }}
Your Ticket is {{.ticket}}</br>
Your username is {{.username}}</br>
Your phone is {{.phone}}
    <a href="http://sso1.xxx.com:8088/logout">本地注销</a>
{{ else  }}
    <div>这只是个演示页面</div>
    <a href="http://sso.xxx.com:8080/sso/login?next=http://sso1.xxx.com:8088">sso server登陆</a>
    <script>
    function ticket_callback(ticket){
        if( ticket != "" ){
            var n = document.createElement("script");
            n.src="/get_ticket?ticket="+ticket;
            document.body.appendChild(n);
            location.reload();
        }
    }
    </script>
    <script src="http://sso.xxx.com:8080/sso/ticket"></script>
{{ end }}

</body>
</html>
