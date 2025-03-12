package renders

const serverError = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Forum | Internal Server Error</title>
</head>
<body>
<div 
    style="display: flex; flex-direction: column; justify-content: center; align-items: center; 
           min-height: 80vh; text-align: center; padding: 20px; background-color: #f8f8f8;"
>
    <h1 style="font-size: 3rem; color: #d9534f; margin: 20px 0;">
        505 - Internal Server Error
    </h1>
    <p style="font-size: 1.2rem; color: #666; margin: 10px 0 30px;">
        Something went wrong on our end. Please try again later.
    </p>
    <a 
        href="/" 
        style="padding: 10px 20px; border: none; background-color: #007bff; color: white; 
               text-decoration: none; border-radius: 5px; transition: background-color 0.3s;"
        onmouseover="this.style.backgroundColor='#0056b3';"
        onmouseout="this.style.backgroundColor='#007bff';"
    >
        Back to Home
    </a>
</div>

</body>
</html>
`