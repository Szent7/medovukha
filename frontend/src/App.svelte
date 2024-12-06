<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import Home from "./routes/Home.svelte";
  import Containers from "./routes/Containers.svelte";

  export let url = "";
</script>

<Router {url}>
  <Route path="/" component={Home} />
  <Route path="/containers" component={Containers} />
</Router>
