<script lang="ts" context="module">
    import { z } from "zod";

    export const ContainerBaseInfoScheme = z.array(
        z.object({
            id: z.string(),
            Names: z.array(z.string()),
            Image: z.string(),
            Ports: z.array(
                z.object({
                    IP: z.string(),
                    PrivatePort: z.number(),
                    PublicPort: z.number(),
                    Type: z.string(),
                }),
            ),
            Created: z.number(),
            State: z.string(),
        }),
    );
    export type ContainerBaseInfo = z.infer<typeof ContainerBaseInfoScheme>;
</script>
