<script lang="ts">
    import FrameElement from "../templates/frameElement.svelte";
    import Sidebar from "../templates/sidebar.svelte";
    import Frame from "../templates/frame.svelte";
    import logo from "../assets/logo_small.svg";
    import { PingBack } from "../lib/api/api.svelte";
</script>

<Sidebar />

<svelte:head>
    <title>Home</title>
    <link rel="icon" type="image/svg+xml" href={logo} />
</svelte:head>
<Frame {content} />

{#snippet header()}
    <img class="logoHeader" src={logo} alt="logo" />
    <p>Home page</p>
    <p>Home page</p>
    <p>Home page</p>
    <p>Home page</p>
    <p>Home page</p>
    <p>Home page</p>
    <button on:click={() => PingBack()}>PING</button>
{/snippet}

{#snippet content()}
    <FrameElement content={header} />
    <FrameElement content={header} />
{/snippet}

<style>
    .logoHeader {
        display: block;
        max-height: 100px;
    }
</style>
