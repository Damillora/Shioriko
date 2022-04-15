<script>
	import { Router, Link, Route } from "svelte-routing";

	import Navbar from "./Navbar.svelte";

	
	import Home from "./routes/Home.svelte";

	import Login from "./routes/Auth/Login.svelte";
	import Logout from "./routes/Auth/Logout.svelte";
	import Register from "./routes/Auth/Register.svelte";

	import Posts from "./routes/Post/Posts.svelte";
	import Post from "./routes/Post/Post.svelte";
	import Upload from "./routes/Post/Upload.svelte";

	import Tags from "./routes/Tags/Tags.svelte";
	import Tag from "./routes/Tags/Tag.svelte";

	import Profile from "./routes/User/Profile.svelte";

	export let url = "";
	let baseURL = window.BASE_URL;
</script>

<Router {url}>
	<Navbar />
	<div>
		<Route path="/" component={Home} />
		<Route path="/posts" component={Posts} />
		<Route path="/post/:id" component={Post} />
		<Route path="/auth/login" component={Login} />
		<Route path="/auth/logout" component={Logout} />
		<Route path="/upload" component={Upload} />
		<Route path="/tags" component={Tags} />
		<Route path="/tags/:tag" component={Tag} />
		<Route path="/auth/register" component={Register} />
		<Route path="/user/profile" component={Profile} />
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
	#tags .svelte-tags-input-matchs {
		z-index: 200;
		&-parent {
			z-index: 200;
		}
		& li:hover {
			background: $primary;
		}
	}
</style>
