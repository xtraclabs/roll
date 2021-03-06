package html

var Authorize = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Authorize Access to XTRAC Data<</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
</head>
<body>



<div class="container">
    <h2>{{.AppName}} Would Like Access to Your XTRAC Data</h2>
<form method="post" role="form" action="validate">
    <div class="form-group">
        <label for="username">User Name:</label>
        <input type="text" class="form-control" id="username" name="username"/>
    </div>
    <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" class="form-control" id="password" name="password"/>
    </div>

    <button type="submit"  class="btn btn-default" name="authorize" value="allow">Allow</button>
    <button type="submit"  class="btn btn-info" name="authorize" value="deny">Deny</button>

    <input type="hidden" name="client_id" value="{{.ClientID}}"/>
    <input type="hidden" name="response_type" value="token"/>
    <input type="hidden" name="scope" value="{{.Scope}}"/>
</form>
</div>
</body>
</html>
`

var Authorize3Leg = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Authorize Access to XTRAC Data<</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
</head>
<body>



<div class="container">
    <h2>{{.AppName}} Would Like Access to Your XTRAC Data</h2>
<form method="post" role="form" action="validate">
    <div class="form-group">
        <label for="username">User Name:</label>
        <input type="text" class="form-control" id="username" name="username"/>
    </div>
    <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" class="form-control" id="password" name="password"/>
    </div>

    <button type="submit"  class="btn btn-default" name="authorize" value="allow">Allow</button>
    <button type="submit"  class="btn btn-info" name="authorize" value="deny">Deny</button>

    <input type="hidden" name="client_id" value="{{.ClientID}}"/>
    <input type="hidden" name="response_type" value="code"/>
    <input type="hidden" name="scope" value="{{.Scope}}"/>
</form>
</div>
</body>
</html>
`
