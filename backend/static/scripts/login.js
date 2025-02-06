document.addEventListener('DOMContentLoaded', function() { 
    // Initialize Lucide icons 
    lucide.createIcons(); 
    // DOM Elements 
    const form = document.getElementById('signinForm'); 
    const emailInput = document.getElementById('email'); 
    const passwordInput = document.getElementById('password'); 
    const passwordToggle = document.getElementById('passwordToggle'); 
    const submitButton = form.querySelector('button[type="submit"]'); 
    const spinner = submitButton.querySelector('.spinner'); 
    // Password visibility toggle 
    passwordToggle.addEventListener('click', () => { 
        const type = passwordInput.type === 'password' ? 'text' : 'password'; passwordInput.type = type; 
        // Update icon 
        const icon = passwordToggle.querySelector('i'); 
        if (icon != null) {
            icon.setAttribute('data-lucide', type === 'password' ? 'eye' : 'eye-off'); 
        }
    }); 

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
        if (error) { 
            error.remove(); 
        } 

        input.classList.remove('error-input'); 
    } 
    
    // Form submission 
    form.addEventListener('submit', async (e) => { 
        e.preventDefault(); 
        // Remove any existing errors 
        removeError(emailInput); 
        removeError(passwordInput); 
        // Validate email 
        if (!validateEmail(emailInput.value)) { 
            showError(emailInput, 'Please enter a valid email address'); 
            return; 
        } 
        
        // Validate password 
        if (passwordInput.value.length < 6) { 
            showError(passwordInput, 'Password must be at least 6 characters'); 
            return; 
        } 
        
        // Show loading state 
        submitButton.disabled = true; 
        spinner.classList.remove('hidden'); 
        submitButton.querySelector('span').textContent = 'Signing in...'; 
        try { 
            // Simulate API call 
            await new Promise(resolve => setTimeout(resolve, 1500)); 
            
            // Redirect to main page (replace with your actual redirect logic) 
            window.location.href = '/'; 
        } catch (error) { 
            showError(emailInput, 'Invalid email or password'); 
            showError(passwordInput, 'Invalid email or password'); 
        } finally { 
            // Reset button state 
            submitButton.disabled = false; spinner.classList.add('hidden'); 
            submitButton.querySelector('span').textContent = 'Sign in'; 
        } 
    }); 
        
    // Clear errors on input 
    emailInput.addEventListener('input', () => removeError(emailInput)); 
    passwordInput.addEventListener('input', () => removeError(passwordInput)); 
});