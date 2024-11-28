<script lang="ts">
    import axios from "axios";

    axios.get("http://localhost:10015/ping").then(function (response) {
        console.log(response.data);
    });
</script>
