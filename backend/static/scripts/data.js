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
};

// Constants
const CONSTANTS = {
	MIN_PASSWORD_LENGTH: 8,
	MAX_PASSWORD_LENGTH: 35,
};

// Posts Data
const POSTS = [];
const COMMENTS = {};
const REPLIES = {};

// User Data
const STATE = {
	currentView: "overview",
	darkMode: localStorage.getItem("darkMode") === "true",
	profilePic: "",
	bio:
		localStorage.getItem("userBio") ||
		"Hi, I love coding and sharing knowledge with the community!",
	posts: [
		{
			id: 1,
			content: "Just learned about React hooks!",
			comments: 5,
			likes: 12,
			timestamp: "2h ago",
		},
		{
			id: 2,
			content: "Working on a new project using TypeScript",
			comments: 3,
			likes: 8,
			timestamp: "5h ago",
		},
		{
			id: 3,
			content: "Check out my latest blog post about web performance",
			comments: 8,
			likes: 15,
			timestamp: "1d ago",
		},
	],
	userComments: [
		{
			id: 1,
			postTitle: "Introduction to GraphQL",
			content: "Great explanation!",
			likes: 5,
			timestamp: "3h ago",
		},
		{
			id: 2,
			postTitle: "Docker Best Practices",
			content: "Very effective!",
			likes: 3,
			timestamp: "1d ago",
		},
	],
	activities: [
		{
			type: "post",
			content: "Created a new post",
			timestamp: "2h ago",
		},
		{
			type: "comment",
			content: "Commented on 'Docker Best Practices'",
			timestamp: "1d ago",
		},
		{
			type: "like",
			content: "Liked 'Introduction to GraphQL'",
			timestamp: "1d ago",
		},
	],
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
};
