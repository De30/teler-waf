package teler

const (
	xTelerReqId  = "X-Teler-Req-Id"
	forbiddenTpl = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>403 Forbidden</title>
</head>
<body>
<center>
	<h1>403 Forbidden</h1>
	We're sorry, but your request has been denied for security reasons.<br>
	If you feel this is an error, please contact customer support for further assistance.
</center>
<hr><center>teler rID: %s</center>
</body>
</html>
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->`
)
