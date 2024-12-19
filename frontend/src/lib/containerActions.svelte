<script lang="ts" context="module">
    import {
        KillContainerById,
        PauseContainerById,
        StartContainerById,
        UnpauseContainerById,
        RemoveContainerById,
    } from "./api/api.svelte";

    export function StartContainer(ids: string[]) {
        ids.forEach((id: string): void => {
            StartContainerById(id);
        });
    }

    export function PauseContainer(ids: string[]) {
        ids.forEach((id: string): void => {
            PauseContainerById(id);
        });
    }

    export function UnpauseContainer(ids: string[]) {
        ids.forEach((id: string): void => {
            UnpauseContainerById(id);
        });
    }

    export function KillContainer(ids: string[]) {
        ids.forEach((id: string): void => {
            KillContainerById(id);
        });
    }

    export function RemoveContainer(ids: string[]) {
        ids.forEach((id: string): void => {
            RemoveContainerById(id);
        });
    }
</script>
