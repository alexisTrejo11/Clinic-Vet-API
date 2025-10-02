package emailTemplates

var MfaTemplate = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Código de Verificación - {{.ProjectName}}</title>
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
        .security-icon {
            text-align: center;
            margin-bottom: 25px;
        }
        .security-icon svg {
            width: 60px;
            height: 60px;
            fill: #667eea;
        }
        .welcome {
            font-size: 18px;
            margin-bottom: 20px;
            color: #2c3e50;
        }
        .mfa-code {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 30px;
            margin: 25px 0;
            text-align: center;
            border-radius: 12px;
            box-shadow: 0 5px 20px rgba(102, 126, 234, 0.3);
        }
        .mfa-code .code {
            font-size: 36px;
            font-weight: bold;
            letter-spacing: 8px;
            font-family: 'Courier New', monospace;
            margin-bottom: 10px;
        }
        .mfa-code .expiry {
            font-size: 14px;
            opacity: 0.9;
        }
        .warning {
            background-color: #fff3cd;
            border-left: 4px solid #ffc107;
            padding: 15px;
            margin: 25px 0;
            border-radius: 4px;
        }
        .warning strong {
            color: #856404;
        }
        .footer {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 25px 30px;
            text-align: center;
            font-size: 14px;
        }
        @media (max-width: 600px) {
            .container {
                margin: 0 10px;
            }
            .header, .content, .footer {
                padding: 20px;
            }
            .mfa-code .code {
                font-size: 28px;
                letter-spacing: 4px;
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
            <p>Verificación de Seguridad</p>
        </div>
        
        <div class="content">
            <div class="security-icon">
                <svg viewBox="0 0 24 24">
                    <path d="M12,1L3,5V11C3,16.55 6.84,21.74 12,23C17.16,21.74 21,16.55 21,11V5L12,1M12,7C13.4,7 14.8,8.6 14.8,10V11.5C15.4,11.5 16,12.4 16,13V16C16,17.4 15.4,18 14.8,18H9.2C8.6,18 8,17.4 8,16V13C8,12.4 8.6,11.5 9.2,11.5V10C9.2,8.6 10.6,7 12,7M12,8.2C11.2,8.2 10.5,8.7 10.5,10V11.5H13.5V10C13.5,8.7 12.8,8.2 12,8.2Z"/>
                </svg>
            </div>
            
            <div class="welcome">¡Hola {{.UserName}}!</div>
            
            <p>Hemos recibido una solicitud para acceder a tu cuenta en <strong>{{.ProjectName}}</strong>. Por tu seguridad, utiliza el siguiente código de verificación:</p>
            
            <div class="mfa-code">
                <div class="code">{{.Token}}</div>
                <div class="expiry">Válido por 10 minutos</div>
            </div>
            
            <div class="warning">
                <strong>Importante:</strong> Si no has intentado iniciar sesión, tu cuenta podría estar en riesgo. Te recomendamos cambiar tu contraseña inmediatamente.
            </div>
            
            <p style="color: #666; font-size: 14px; margin-top: 30px;">
                Este código es personal e intransferible. No lo compartas con nadie.
            </p>
        </div>
        
        <div class="footer">
            <p>&copy; {{.Year}} {{.ProjectName}}. Todos los derechos reservados.</p>
            <p>Este es un correo automático, por favor no respondas a esta dirección.</p>
        </div>
    </div>
</body>
</html>
`
