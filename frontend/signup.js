document.addEventListener('DOMContentLoaded', function() { 
    // Initialize Lucide icons 
    lucide.createIcons(); 

    // DOM Elements 
    const form = document.getElementById('signupForm'); 
    const usernameInput = document.getElementById('username'); 
    const emailInput = document.getElementById('email'); 
    const passwordInput = document.getElementById('password'); 
    const confirmPasswordInput = document.getElementById('confirmPassword'); 
    const termsCheckbox = document.getElementById('terms'); 
    const passwordToggle = document.getElementById('passwordToggle'); 
    const strengthBars = document.querySelectorAll('.strength-bars .bar'); 
    const strengthText = document.querySelector('.strength-text'); 
    const submitButton = form.querySelector('button[type="submit"]'); 
    const spinner = submitButton.querySelector('.spinner'); 

    // Password visibility toggle 
    passwordToggle.addEventListener('click', () => { 
        const type = passwordInput.type === 'password' ? 'text' : 'password'; 
        passwordInput.type = type; 
        confirmPasswordInput.type = type; 
        // Update icon 
        const icon = passwordToggle.querySelector('i'); 
        icon.setAttribute('data-lucide', type === 'password' ? 'eye' : 'eye-off'); 
        lucide.createIcons(); 
    }); 

    // Password strength checker 
    function checkPasswordStrength(password) { 
        let strength = 0; 
        if (password.length >= 8) strength++; 
        if (password.match(/[a-z]/) && password.match(/[A-Z]/)) strength++; 
        if (password.match(/\d/)) strength++; if (password.match(/[^a-zA-Z\d]/)) strength++; 
        
        // Update strength bars 
        strengthBars.forEach((bar, index) => { 
            bar.className = 'bar'; 
            if (index < strength) { 
                bar.classList.add(strength <= 2 ? 'weak' : strength === 3 ? 'medium' : 'strong'); 
            } 
        }); 

        // Update strength text 
        const strengthLabels = ['Weak', 'Fair', 'Good', 'Strong']; 
        strengthText.textContent = strength > 0 ? strengthLabels[strength - 1] : 'Password strength'; 
        return strength; 
    } 
});