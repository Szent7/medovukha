<script lang="ts">
    import FrameElement from "../templates/frameElement.svelte";
    import Sidebar from "../templates/sidebar.svelte";
    import Frame from "../templates/frame.svelte";
    import logo from "../assets/logo.svg";
    import { PingBack } from "../lib/api/api.svelte";
</script>

<Sidebar />
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
