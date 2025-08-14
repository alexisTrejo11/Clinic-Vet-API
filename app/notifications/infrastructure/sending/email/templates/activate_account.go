package emailTemplates

var ActivationTemplate = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Activación de Cuenta - {{.ProjectName}}</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: #333;
            background-color: #f4f7fa;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 30px;
            text-align: center;
        }
        .header img {
            max-width: 120px;
            height: auto;
            margin-bottom: 15px;
        }
        .header h1 {
            font-size: 28px;
            font-weight: 300;
            margin-bottom: 5px;
        }
        .header p {
            font-size: 16px;
            opacity: 0.9;
        }
        .content {
            padding: 40px 30px;
        }
        .welcome {
            font-size: 18px;
            margin-bottom: 20px;
            color: #2c3e50;
        }
        .token-box {
            background: #f8f9ff;
            border: 2px dashed #667eea;
            padding: 25px;
            margin: 25px 0;
            text-align: center;
            border-radius: 8px;
        }
        .token {
            font-size: 32px;
            font-weight: bold;
            color: #667eea;
            letter-spacing: 3px;
            font-family: 'Courier New', monospace;
        }
        .button {
            display: inline-block;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 15px 30px;
            text-decoration: none;
            border-radius: 25px;
            font-weight: 600;
            margin: 20px 0;
            transition: all 0.3s ease;
        }
        .button:hover {
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(102, 126, 234, 0.3);
        }
        .footer {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 25px 30px;
            text-align: center;
            font-size: 14px;
        }
        .footer p {
            margin-bottom: 10px;
        }
        .social-links {
            margin-top: 15px;
        }
        .social-links a {
            color: #3498db;
            text-decoration: none;
            margin: 0 10px;
        }
        @media (max-width: 600px) {
            .container {
                margin: 0 10px;
            }
            .header, .content, .footer {
                padding: 20px;
            }
            .token {
                font-size: 24px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            {{if .LogoURL}}
            <img src="{{.LogoURL}}" alt="{{.ProjectName}} Logo">
            {{end}}
            <h1>{{.ProjectName}}</h1>
            <p>Activación de Cuenta</p>
        </div>
        
        <div class="content">
            <div class="welcome">¡Hola {{.UserName}}!</div>
            
            <p>Gracias por registrarte en <strong>{{.ProjectName}}</strong>. Para completar tu registro y activar tu cuenta, utiliza el siguiente código de activación:</p>
            
            <div class="token-box">
                <div class="token">{{.Token}}</div>
                <p style="margin-top: 10px; color: #666;">Código de activación</p>
            </div>
            
            <p>Este código es válido por 24 horas. Si no has solicitado esta activación, puedes ignorar este correo.</p>
            
            {{if .ButtonURL}}
            <div style="text-align: center; margin: 30px 0;">
                <a href="{{.ButtonURL}}" class="button">{{.ButtonText}}</a>
            </div>
            {{end}}
            
            <p style="color: #666; font-size: 14px; margin-top: 30px;">
                Si tienes alguna pregunta, no dudes en contactarnos. ¡Esperamos verte pronto!
            </p>
        </div>
        
        <div class="footer">
            <p>&copy; {{.Year}} {{.ProjectName}}. Todos los derechos reservados.</p>
            <p>Este es un correo automático, por favor no respondas a esta dirección.</p>
        </div>
    </div>
</body>
</html>`
