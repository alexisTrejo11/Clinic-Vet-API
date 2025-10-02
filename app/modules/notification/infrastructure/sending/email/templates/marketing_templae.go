package emailTemplates

var MarketingTemplate = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} - {{.ProjectName}}</title>
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
        .title {
            font-size: 24px;
            color: #2c3e50;
            margin-bottom: 20px;
            text-align: center;
        }
        .welcome {
            font-size: 18px;
            margin-bottom: 25px;
            color: #2c3e50;
        }
        .message-content {
            background-color: #f8f9ff;
            padding: 25px;
            border-radius: 8px;
            margin: 25px 0;
            border-left: 4px solid #667eea;
        }
        .cta-section {
            text-align: center;
            margin: 35px 0;
        }
        .button {
            display: inline-block;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 16px 35px;
            text-decoration: none;
            border-radius: 30px;
            font-weight: 600;
            font-size: 16px;
            transition: all 0.3s ease;
            box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
        }
        .button:hover {
            transform: translateY(-3px);
            box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
        }
        .features {
            margin: 30px 0;
        }
        .feature-item {
            display: flex;
            align-items: center;
            margin-bottom: 15px;
            padding: 15px;
            background-color: #f8f9fa;
            border-radius: 8px;
        }
        .feature-icon {
            width: 20px;
            height: 20px;
            margin-right: 15px;
            fill: #667eea;
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
        .unsubscribe {
            margin-top: 15px;
            font-size: 12px;
            opacity: 0.8;
        }
        .unsubscribe a {
            color: #3498db;
            text-decoration: none;
        }
        @media (max-width: 600px) {
            .container {
                margin: 0 10px;
            }
            .header, .content, .footer {
                padding: 20px;
            }
            .title {
                font-size: 20px;
            }
            .feature-item {
                flex-direction: column;
                text-align: center;
            }
            .feature-icon {
                margin-right: 0;
                margin-bottom: 10px;
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
            <p>{{.Title}}</p>
        </div>
        
        <div class="content">
            <div class="title">{{.Title}}</div>
            
            <div class="welcome">¡Hola {{.UserName}}!</div>
            
            <div class="message-content">
                {{.Message}}
            </div>
            
            {{if .ButtonURL}}
            <div class="cta-section">
                <a href="{{.ButtonURL}}" class="button">{{.ButtonText}}</a>
            </div>
            {{end}}
            
            <div class="features">
                <div class="feature-item">
                    <svg class="feature-icon" viewBox="0 0 24 24">
                        <path d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z"/>
                    </svg>
                    <span>Servicio de alta calidad garantizado</span>
                </div>
                <div class="feature-item">
                    <svg class="feature-icon" viewBox="0 0 24 24">
                        <path d="M12,1L3,5V11C3,16.55 6.84,21.74 12,23C17.16,21.74 21,16.55 21,11V5L12,1Z"/>
                    </svg>
                    <span>Información segura y protegida</span>
                </div>
                <div class="feature-item">
                    <svg class="feature-icon" viewBox="0 0 24 24">
                        <path d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M16.2,16.2L11,13V7H12.5V12.2L17,14.9L16.2,16.2Z"/>
                    </svg>
                    <span>Atención 24/7 cuando nos necesites</span>
                </div>
            </div>
            
            <p style="color: #666; font-size: 14px; margin-top: 30px;">
                Gracias por confiar en nosotros. ¡Esperamos poder ayudarte pronto!
            </p>
        </div>
        
        <div class="footer">
            <p>&copy; {{.Year}} {{.ProjectName}}. Todos los derechos reservados.</p>
            <div class="unsubscribe">
                <p>¿No deseas recibir estos correos? <a href="#">Darse de baja</a></p>
            </div>
        </div>
    </div>
</body>
</html>`
