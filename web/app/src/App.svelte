<script>
	import { Router, Link, Route } from "svelte-routing";

	import Navbar from "./Navbar.svelte";

	import Home from "./routes/Home.svelte";
	import Posts from "./routes/Posts.svelte";
	import Post from "./routes/Post.svelte";
	import Login from "./routes/Login.svelte";
	import Logout from "./routes/Logout.svelte";
	import Upload from "./routes/Upload.svelte";
	import Edit from "./routes/Edit.svelte";

	export let url = "";
	let baseURL = window.BASE_URL;
</script>

<Router {url}>
	<Navbar />
	<div>
		<Route path="/" component={Home} />
		<Route path="/posts" component={Posts} />
		<Route path="/post/:id" component={Post} />
		<Route path="/post/edit/:id" component={Edit} />
		<Route path="/auth/login" component={Login} />
		<Route path="/auth/logout" component={Logout} />
		<Route path="/upload" component={Upload} />
	</div>
</Router>

<style global lang="scss">
	@import "./main.scss";

	#tags .svelte-tags-input-tag {
		background: $primary;
		color: $text-invert;
	}

	#tags .svelte-tags-input-layout {
		@extend .input;
		height: inherit;

		& .svelte-tags-input {
			margin-top: 0 !important;
			font-size: 13.3333px;
		}
	}
</style>
