<div class="sidebar">
	<h3 id="medovukhaTitle">Medovukha</h3>
	<div id="navMenu" class="menuItem">
		<nav>
			<li><a href="/">Home</a></li>
			<li><a href="/containers">Containers</a></li>
			<li><a href="/images">Images</a></li>
			<li><a href="/networks">Networks</a></li>
			<li><a href="/volumes">Volumes</a></li>
		</nav>
	</div>
	<!-- <div class="menuItem" id="version">0.0.1</div>-->
</div>

<style>
	#medovukhaTitle {
		margin: 25px auto 25px auto;
		text-align: center;
		font-size: 20pt;
		font-variant: small-caps;
	}

	/*#version {
        text-align: center;
        margin: auto 55px 25px 55px;
    }*/

	li {
		text-align: center;
	}

	.menuItem {
		background-color: #1c1c1c;
		margin: 0 15px 0 15px;
		border-radius: 10px;
	}

	:global(a) {
		color: #e3e3e3;
		display: block;
		border-width: 2px;
		border-color: transparent;
		border-radius: 10px;
		border-style: solid;
		box-sizing: border-box;
		transition:
			border-color 0.25s,
			background-color 0.25s;
	}

	:global(a):hover {
		border-color: aqua;
		color: aqua;
		background-color: #33f8ff4d;
	}

	li {
		list-style-type: none;
	}

	.sidebar {
		width: 250px;
		height: 100%;
		background-color: #2d2d2d;
		-webkit-box-shadow: 0px 10px 10px 0px rgba(0, 0, 0, 0.2);
		-moz-box-shadow: 0px 10px 10px 0px rgba(0, 0, 0, 0.2);
		box-shadow: 0px 10px 10px 0px rgba(0, 0, 0, 0.2);
		display: flex;
		flex-direction: column;
		float: inline-start;
	}
</style>
