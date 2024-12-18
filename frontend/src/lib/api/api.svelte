<script lang="ts" context="module">
    import axios from "axios";
    import {
        ContainerBaseInfoScheme,
        type ContainerBaseInfo,
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
</script>
