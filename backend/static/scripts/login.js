document.addEventListener('DOMContentLoaded', function () {
    // Initialize Lucide icons
    lucide.createIcons();

    // DOM Elements
    const form = document.getElementById('signinForm');
    const emailInput = document.getElementById('email');
    const passwordInput = document.getElementById('password');
    const passwordToggle = document.getElementById('passwordToggle');
    const submitButton = form.querySelector('button[type="submit"]');
    const spinner = submitButton.querySelector('.spinner');

    // Constants
    const API_ENDPOINT = '/api/auth/login';
    const MIN_PASSWORD_LENGTH = 6;

    // Password visibility toggle
    passwordToggle.addEventListener('click', togglePasswordVisibility);

    function togglePasswordVisibility() {
        const isPassword = passwordInput.type === 'password';
        passwordInput.type = isPassword ? 'text' : 'password';

        if (passwordToggle.querySelector('i')) {
            passwordToggle.querySelector('i').setAttribute('data-lucide', isPassword ? 'eye-off' : 'eye');
            lucide.createIcons(passwordToggle);
        }
    }

    // Form validation
    function validateEmail(email) {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    }

    function showError(input, message) {
        const formGroup = input.closest('.form-group');
        let errorDiv = formGroup.querySelector('.error');

        if (!errorDiv) {
            errorDiv = document.createElement('div');
            errorDiv.className = 'error';
            formGroup.appendChild(errorDiv);
        }

        errorDiv.textContent = message;
        input.classList.add('error-input');
    }

    function removeError(input) {
        const formGroup = input.closest('.form-group');
        const error = formGroup.querySelector('.error');
        if (error) error.remove();
        input.classList.remove('error-input');
    }

    // Form submission
    form.addEventListener('submit', handleLogin);

    async function handleLogin(e) {
        e.preventDefault();

        removeError(emailInput);
        removeError(passwordInput);

        if (!validateEmail(emailInput.value)) {
            showError(emailInput, 'Please enter a valid email address');
            return;
        }

        if (passwordInput.value.length < MIN_PASSWORD_LENGTH) {
            showError(passwordInput, `Password must be at least ${MIN_PASSWORD_LENGTH} characters`);
            return;
        }

        // Show loading state
        submitButton.disabled = true;
        spinner.classList.remove('hidden');
        submitButton.querySelector('span').textContent = 'Signing in...';

        const formData = {
            email: emailInput.value,
            password: passwordInput.value,
        };

        try {
            const response = await fetch(API_ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(formData),
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            const res = await response.json();

            if (response.ok) {
                console.log('Login successful:', res.message);
                localStorage.setItem('res', JSON.stringify(res.data));
                window.location.href = '/';
            } else {
                console.error('Login failed:', res.message || 'Unknown error');
                showError(emailInput, res.message || 'Invalid email or password');
            }
        } catch (error) {
            console.error('Login Error:', error.message);
            showError(emailInput, 'An error occurred. Please try again.');
        } finally {
            submitButton.disabled = false;
            spinner.classList.add('hidden');
            submitButton.querySelector('span').textContent = 'Sign in';
        }
    }

    // Clear errors dynamically
    form.addEventListener('input', (event) => {
        if (event.target.matches('#email, #password')) {
            removeError(event.target);
        }
    });
});
