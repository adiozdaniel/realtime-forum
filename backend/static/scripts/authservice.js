import { API_ENDPOINTS } from "./data.js";

// AuthService class for handling authentication-related API requests
class AuthService {
	constructor() {
		this.apiEndpoints = API_ENDPOINTS;
	}
}

// Method to log in a user
AuthService.prototype.login = async function (credentials) {
	if (!credentials?.email || !credentials?.password) {
		return (data = {
			error: true,
			message: "Please provide both email and password!",
		});
	}

	try {
		const response = await fetch(this.apiEndpoints.login, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(credentials),
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

// Method to register a new user
AuthService.prototype.register = async function (formData) {
	if (!formData?.email || !formData?.password || !formData?.user_name) {
		return (data = {
			error: true,
			message: "Please provide all required fields!",
		});
	}

	try {
		const response = await fetch(this.apiEndpoints.register, {
			method: "POST",
			body: JSON.stringify(formData),
			headers: {
				"Content-Type": "application/json",
			},
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

// Method to log out a user
AuthService.prototype.logout = async function () {
	try {
		const response = await fetch(this.apiEndpoints.logout, {
			method: "POST",
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

// Method to check if the user is authenticated
AuthService.prototype.isAuthenticated = async function () {
	try {
		const response = await fetch(this.apiEndpoints.check, {
			method: "GET",
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

// Method to upload a profile picture
AuthService.prototype.uploadProfilePic = async function (formData) {
	try {
		const response = await fetch(this.apiEndpoints.uploadProfilePic, {
			method: "POST",
			body: formData,
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

export { AuthService };
