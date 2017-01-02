<html>
    <head>
        <title>Registration</title>
    </head>
    <body>
       
        <form action="/register" method="post">
            <table align="center">
                
                <tr>
                    <td>Nome:</td>
                    <td><input type="text" name="name"></td>
                </tr>
                
                <tr>
                    <td>Cognome:</td>
                    <td><input type="text" name="cognome"></td>
                </tr>            
                
                <tr>
                    <td>e-mail:</td>
                    <td><input type="text" name="email"></td>                
                </tr>
                
                <tr>
                    <td>Username:</td>
                    <td><input type="text" name="username"></td>               
                </tr>

                <tr>
                    <td>Password:</td>
                    <td><input type="password" name="password"></td>
                </tr>
                
                <tr>
                    <td></td>
                    <td><input type="submit" value="Register"></td>
                </tr>
            </table>
        </form>
    </body>
</html>