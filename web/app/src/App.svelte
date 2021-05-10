<script>
	import { Router, Link, Route } from "svelte-routing";

	import Navbar from "./Navbar.svelte";

	import Home from "./routes/Home.svelte";
	import Posts from "./routes/Posts.svelte";
	import Post from "./routes/Post.svelte";
	import Login from "./routes/Login.svelte";
	import Logout from "./routes/Logout.svelte";
	import Tag from "./routes/Tag.svelte";

	

	export let url = "";
	let baseURL = window.BASE_URL;

</script>

<Router {url}>
	<Navbar />
	<div>
		<Route path="/" component={Home} />
		<Route path="/posts" component={Posts} />
		<Route path="/tag/:id" component={Tag} />
		<Route path="/post/:id" component={Post} />
		<Route path="/auth/login" component={Login} />
		<Route path="/auth/logout" component={Logout} />
	</div>
</Router>
