// API Endpoints
const API_ENDPOINTS = {
	login: "/api/auth/login",
	register: "/api/auth/register",
	logout: "/api/auth/logout",
	check: "/api/auth/check",
	uploadProfilePic: "/api/auth/uploadProfilePic",
	editBio: "/api/user/editBio",
	userDashBoard: "/api/user/dashboard",

	// posts ENDPOINTS
	allposts: "/api/posts",
	createpost: "/api/posts/create",
	deletepost: "/api/posts/delete",
	updatepost: "/api/posts/update",
	likepost: "/api/posts/like",
	dislikepost: "/api/posts/dislike",
	uploadPostImg: "/api/posts/image",

	// comments ENDPOINTS
	listcommbypost: "/api/comments",
	createcomment: "/api/posts/comments/create",
	deletecomment: "/api/comments/delete",
	updatecomment: "/api/comments/update",
	likecomment: "/api/comments/like",
	dislikecomment: "/api/comments/dislike",

	// reply ENDPOINTS
	createReply: "/api/comments/reply/create",

	// notification ENDPOINTS
	checkNotifications: "/api/notifications/check",
	readNotifications: "/api/notifications/read"
};

// Constants
const CONSTANTS = {
	MIN_PASSWORD_LENGTH: 8,
	MAX_PASSWORD_LENGTH: 35,
};

// Temp Data
let TEMP_DATA = null;

// Posts Data
const POSTS = [];
const COMMENTS = {};
const REPLIES = {};

// User Data
const STATE = {
	currentView: "overview",
	darkMode: "",
	profilePic: "",
	bio: "Hi, I love coding and sharing knowledge with the community!",
	posts: [],
	userComments: [],
	activities: [],
};

const commentLikeState = {
	comments: {},
};

const postLikeState = {
	posts: {},
	comments: {},
};

const postDislikeState = {
	posts: {},
	comments: {},
};

export {
	API_ENDPOINTS,
	POSTS,
	CONSTANTS,
	STATE,
	commentLikeState,
	postLikeState,
	postDislikeState,
	COMMENTS,
	REPLIES,
	TEMP_DATA,
};
