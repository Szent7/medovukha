<script lang="ts">
    import { onMount } from "svelte";
    import FrameElement from "../templates/frameElement.svelte";
    import Sidebar from "../templates/sidebar.svelte";
    import Frame from "../templates/frame.svelte";
    import logo from "../assets/logo.svg";
    import { GetContainerList } from "../lib/api/api.svelte";
    import type { ContainerBaseInfo } from "../lib/api/types.svelte";
    import { UnixTimeFormat } from "../lib/time.svelte";
    import {
        KillContainer,
        PauseContainer,
        RemoveContainer,
        StartContainer,
        UnpauseContainer,
    } from "../lib/containerActions.svelte";

    let conList: ContainerBaseInfo = [];
    let loading = true;
    let selectedIds: string[] = [];
    const buttonIds: string[] = [
        "start-button",
        "stop-button",
        "kill-button",
        "restart-button",
        "pause-button",
        "resume-button",
        "remove-button",
    ];
    let activateButtons: boolean = false;

    const updateContainerList = () => {
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
    };

    onMount(updateContainerList);

    function updateSelected(id: string, checked: boolean) {
        if (checked) {
            selectedIds = [...selectedIds, id];
            if (!activateButtons) {
                buttonIds.forEach((item: string): void => {
                    const button = document.getElementById(
                        item,
                    ) as HTMLButtonElement | null;
                    if (button !== null) {
                        button.disabled = false;
                    }
                });
                activateButtons = true;
            }
        } else {
            selectedIds = selectedIds.filter((item) => item !== id);
            if (selectedIds.length == 0) {
                buttonIds.forEach((item: string): void => {
                    const button = document.getElementById(
                        item,
                    ) as HTMLButtonElement | null;
                    if (button !== null) {
                        button.disabled = true;
                    }
                });
                activateButtons = false;
            }
        }
        console.log("SelectedIds:" + selectedIds);
    }

    function selectAll(checked: boolean) {
        const elements = document.querySelectorAll("input[name=checkbox-item]");
        if (elements !== null) {
            Array.prototype.forEach.call(elements, function (item) {
                if (item.checked != checked) {
                    updateSelected(item.id, checked);
                    item.checked = checked;
                }
            });
        }
    }
</script>

<Sidebar />
<Frame {content} />

{#snippet content()}
    <FrameElement content={header} />
    {#if loading}
        <FrameElement content={loadingData} />
    {:else if !loading}
        <FrameElement content={containerTable} />
    {/if}
{/snippet}

{#snippet loadingData()}
    <p>Loading data...</p>
{/snippet}

{#snippet header()}
    <img class="logoHeader" src={logo} alt="logo" />
    <p>Container page</p>
{/snippet}

{#snippet containerTable()}
    <div class="button-block">
        <button
            class="containerlist-button"
            id="start-button"
            onclick={() => {
                StartContainer(selectedIds);
                setTimeout(() => {
                    updateContainerList();
                }, 1000);
            }}
            disabled>Start</button
        >
        <button
            class="containerlist-button"
            id="stop-button"
            onclick={() => {
                //StopContainer(selectedIds);
            }}
            disabled>Stop</button
        >
        <button
            class="containerlist-button"
            id="kill-button"
            onclick={() => {
                KillContainer(selectedIds);
                setTimeout(() => {
                    updateContainerList();
                }, 1000);
            }}
            disabled>Kill</button
        >
        <button
            class="containerlist-button"
            id="restart-button"
            onclick={() => {
                //RestartContainer(selectedIds);
            }}
            disabled>Restart</button
        >
        <button
            class="containerlist-button"
            id="pause-button"
            onclick={() => {
                PauseContainer(selectedIds);
                setTimeout(() => {
                    updateContainerList();
                }, 1000);
            }}
            disabled>Pause</button
        >
        <button
            class="containerlist-button"
            id="resume-button"
            onclick={() => {
                UnpauseContainer(selectedIds);
                setTimeout(() => {
                    updateContainerList();
                }, 1000);
            }}
            disabled>Resume</button
        >
        <button
            class="containerlist-button"
            id="remove-button"
            onclick={() => {
                RemoveContainer(selectedIds);
                setTimeout(() => {
                    updateContainerList();
                }, 1000);
            }}
            disabled>Remove</button
        >
    </div>
    <table class="containerlist-table">
        <thead>
            <tr>
                <th
                    ><input
                        type="checkbox"
                        onchange={(event) => {
                            const target = event.target as HTMLInputElement;
                            selectAll(target.checked);
                        }}
                    /></th
                >
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
                    <td
                        ><input
                            type="checkbox"
                            name="checkbox-item"
                            id={container.id}
                            onchange={(event) => {
                                const target = event.target as HTMLInputElement;
                                updateSelected(container.id, target.checked);
                            }}
                        /></td
                    >
                    <td>{container.Names[0]}</td>
                    <td class="state-{container.State}">{container.State}</td>
                    <td>{container.Image}</td>
                    <td
                        >{#if container.Ports.length == 0}
                            -
                        {:else}
                            {#each container.Ports as port}
                                {#if port?.PublicPort !== undefined && port?.PrivatePort !== undefined}
                                    <p>{port.PublicPort}:{port.PrivatePort}</p>
                                {:else}
                                    -
                                {/if}
                            {/each}
                        {/if}</td
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

    .containerlist-button {
        border-radius: 10px;
    }

    .button-block {
        text-align: right;
    }

    .state-running {
        color: aquamarine;
    }

    .state-exited {
        color: red;
    }

    .state-paused {
        color: orange;
    }
</style>
