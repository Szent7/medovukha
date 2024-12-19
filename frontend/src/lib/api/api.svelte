<script lang="ts" context="module">
    import axios from "axios";
    import {
        ContainerBaseInfoScheme,
        type ContainerBaseInfo,
        type ContainerIdMessage,
    } from "./types.svelte";

    export async function PingBack() {
        axios.get("http://localhost:10015/ping").then(function (response) {
            console.log(response.data);
        });
    }

    export async function GetContainerList() {
        try {
            const response = await axios.get(
                "http://localhost:10015/rest/v1/getcontainerlist",
            );
            const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;
        } catch (error) {
            console.error("Error GET ContainerList:", error);
        }
    }

    export async function StartContainerById(id: string) {
        const container: ContainerIdMessage = {
            Id: id,
        };
        try {
            const response = await axios.post(
                "http://localhost:10015/rest/v1/startcontainerbyid",
                container,
            );
            /*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
        } catch (error) {
            console.error("Error POST StartContainerById:", error);
        }
    }

    export async function PauseContainerById(id: string) {
        const container: ContainerIdMessage = {
            Id: id,
        };
        try {
            const response = await axios.post(
                "http://localhost:10015/rest/v1/pausecontainerbyid",
                container,
            );
            /*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
        } catch (error) {
            console.error("Error POST StartContainerById:", error);
        }
    }

    export async function UnpauseContainerById(id: string) {
        const container: ContainerIdMessage = {
            Id: id,
        };
        try {
            const response = await axios.post(
                "http://localhost:10015/rest/v1/unpausecontainerbyid",
                container,
            );
            /*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
        } catch (error) {
            console.error("Error POST StartContainerById:", error);
        }
    }

    export async function KillContainerById(id: string) {
        const container: ContainerIdMessage = {
            Id: id,
        };
        try {
            const response = await axios.post(
                "http://localhost:10015/rest/v1/killcontainerbyid",
                container,
            );
            /*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
        } catch (error) {
            console.error("Error POST StartContainerById:", error);
        }
    }

    export async function RemoveContainerById(id: string) {
        const container: ContainerIdMessage = {
            Id: id,
        };
        try {
            const response = await axios.post(
                "http://localhost:10015/rest/v1/removecontainerbyid",
                container,
            );
            /*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
        } catch (error) {
            console.error("Error POST StartContainerById:", error);
        }
    }
</script>
