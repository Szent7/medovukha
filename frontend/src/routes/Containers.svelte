<script lang="ts">
    import { onMount } from "svelte";
    import FrameElement from "../templates/frameElement.svelte";
    import Sidebar from "../templates/sidebar.svelte";
    import Frame from "../templates/frame.svelte";
    import logo from "../assets/logo.svg";
    import { GetContainerList } from "../lib/api/api.svelte";
    import type { ContainerBaseInfo } from "../lib/api/types.svelte";
    import { UnixTimeFormat } from "../lib/time.svelte";

    let conList: ContainerBaseInfo = [];
    let loading = true;

    onMount(() => {
        const getData = async () => {
            try {
                const data = await GetContainerList();
                if (data) {
                    conList = data;
                    console.log("data:" + data);
                }
            } catch (err) {
                console.log(err);
            } finally {
                loading = false;
                console.log(loading + "\nconList:" + JSON.stringify(conList));
            }
        };
        getData();
    });
</script>

<Sidebar />
<Frame {content} />

{#snippet content()}
    <FrameElement content={header} />
    {#if loading}
        <p>Loading data...</p>
    {:else if !loading}
        <FrameElement content={containerTable} />
    {/if}
    <!--<FrameElement content={containerTable} />-->
{/snippet}

{#snippet header()}
    <img class="logoHeader" src={logo} alt="logo" />
    <p>Container page</p>
{/snippet}

{#snippet containerTable()}
    <table class="containerlist-table">
        <thead>
            <tr>
                <th><input type="checkbox" /></th>
                <th>Name</th>
                <th>State</th>
                <th>Image</th>
                <th>Port bindings</th>
                <th>Created</th>
            </tr>
        </thead>
        <tbody>
            {#each conList as container}
                <tr>
                    <td><input type="checkbox" /></td>
                    <td>{container.Names[0]}</td>
                    <td>{container.State}</td>
                    <td>{container.Image}</td>
                    <td
                        >{container.Ports[0].PublicPort}:{container.Ports[0]
                            .PrivatePort}</td
                    >
                    <td>{UnixTimeFormat(container.Created)}</td>
                </tr>
            {/each}
        </tbody>
    </table>
{/snippet}

<style>
    .logoHeader {
        display: block;
        max-height: 100px;
    }

    .containerlist-table {
        width: 100%;
        text-align: center;
    }
</style>
