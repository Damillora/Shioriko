<script>
	import { Router, Link, Route } from "svelte-routing";
	import Home from "./routes/Home.svelte";
	import Posts from "./routes/Posts.svelte";
	import Post from "./routes/Post.svelte";
	import Login from "./routes/Login.svelte";
	import Logout from "./routes/Logout.svelte";
	import Tag from "./routes/Tag.svelte";

	import { token } from "./stores.js"
	
	let loggedIn = false;
	token.subscribe(value => {
		loggedIn = value !== ""
	});

	export let url = "";
	let baseURL = window.BASE_URL;

</script>

<Router url="{url}">
	<nav class="navbar" role="navigation" aria-label="main navigation">
		<div class="navbar-brand">
		  <Link class="navbar-item" to="/">
			Shioriko
		  </Link>
	  
		  <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
			<span aria-hidden="true"></span>
			<span aria-hidden="true"></span>
			<span aria-hidden="true"></span>
		  </a>
		</div>
	  
		<div id="navbarBasicExample" class="navbar-menu">
		  <div class="navbar-start">
			<Link class="navbar-item" to="/posts">
			Posts
			</Link>
		  </div>
	  
		  <div class="navbar-end">
			  {#if loggedIn}
			  <div class="navbar-item">
				<div class="buttons">
				  <Link to="/auth/logout" class="button is-light">
					Log out
				  </Link>
				</div>
			  </div>
			  {:else}
			  <div class="navbar-item">
				<div class="buttons">
				  <Link to="/auth/register" class="button is-primary">
					<strong>Register</strong>
				  </Link>
				  <Link to="/auth/login" class="button is-light">
					Log in
				  </Link>
				</div>
			  </div>
			  {/if}
		  </div>
		</div>
	  </nav>
	  <div>
		  <Route path="/" component="{Home}" />
		  <Route path="/posts" component="{Posts}" />
		  <Route path="/tag/:id" component="{Tag}" />
		  <Route path="/post/:id" component="{Post}" />
		  <Route path="/auth/login" component="{Login}" />
		  <Route path="/auth/logout" component="{Logout}" />
	  </div> 
</Router>