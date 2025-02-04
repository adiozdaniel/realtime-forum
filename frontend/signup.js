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

    // Form validation 
    function validateEmail(email) { 
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email); 
    } 

    function showError(input, message) { 
        const formGroup = input.closest('.form-group'); 
        const existingError = formGroup.querySelector('.error'); 
        if (!existingError) { 
            const errorDiv = document.createElement('div'); 
            errorDiv.className = 'error'; 
            errorDiv.textContent = message; 
            formGroup.appendChild(errorDiv); 
        } 
        input.classList.add('error-input'); 
    } 

    function removeError(input) { 
        const formGroup = input.closest('.form-group'); 
        const error = formGroup.querySelector('.error'); 
        if (error) { error.remove(); 

        } 
        input.classList.remove('error-input'); 
    } 

    // Input event listeners 
    passwordInput.addEventListener('input', () => { 
        removeError(passwordInput); 
        checkPasswordStrength(passwordInput.value); 
    }); 

    confirmPasswordInput.addEventListener('input', () => { 
        removeError(confirmPasswordInput); 
        if (confirmPasswordInput.value && confirmPasswordInput.value !== passwordInput.value) { 
            showError(confirmPasswordInput, 'Passwords do not match'); 
        } 
    }); 

    // Form submission 
    form.addEventListener('submit', async (e) => { 
        e.preventDefault(); let isValid = true; 

        // Clear previous errors 
        [usernameInput, emailInput, passwordInput, confirmPasswordInput].forEach(removeError); 

        // Validate username 
        if (usernameInput.value.length < 3) { 
            showError(usernameInput, 'Username must be at least 3 characters'); 
            isValid = false; 
        } 

        // Validate email 
        if (!validateEmail(emailInput.value)) { 
            showError(emailInput, 'Please enter a valid email address'); 
            isValid = false; 
        } 

        // Validate password 
        const passwordStrength = checkPasswordStrength(passwordInput.value); if (passwordStrength < 3) { 
            showError(passwordInput, 'Password is too weak'); 
            isValid = false; 
        } 

        // Validate password confirmation 
        if (passwordInput.value !== confirmPasswordInput.value) { 
            showError(confirmPasswordInput, 'Passwords do not match'); 
            isValid = false; 
        } 

        // Validate terms 
        if (!termsCheckbox.checked) { 
            showError(termsCheckbox, 'You must accept the terms and conditions'); 
            isValid = false; 
        } 

        if (!isValid) return; 

        // Show loading state 
        submitButton.disabled = true; 
        spinner.classList.remove('hidden'); 
        submitButton.querySelector('span').textContent = 'Creating account...'; 
        
        try { 
            // Simulate API call 
            await new Promise(resolve => setTimeout(resolve, 1500)); 
            // Redirect to main page (replace with your actual redirect logic) 
            window.location.href = '/'; 
        } catch (error) { 
            showError(emailInput, 'Failed to create account. Please try again.'); 
        } finally { 
            // Reset button state 
            submitButton.disabled = false; spinner.classList.add('hidden'); submitButton.querySelector('span').textContent = 'Create account'; 
        } 
    }); 
});