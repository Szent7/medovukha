<script lang="ts" context="module">
    import { z } from "zod";

    export const ContainerBaseInfoScheme = z.array(
        z.object({
            id: z.string(),
            Names: z.array(z.string()),
            Image: z.string(),
            Ports: z.array(
                z
                    .object({
                        IP: z.string(),
                        PrivatePort: z.number(),
                        PublicPort: z.number(),
                        Type: z.string(),
                    })
                    .optional(),
            ),
            Created: z.number(),
            State: z.string(),
        }),
    );
    export type ContainerBaseInfo = z.infer<typeof ContainerBaseInfoScheme>;

    export const ContainerIdMessageScheme = z.object({
        Id: z.string(),
    });
    export type ContainerIdMessage = z.infer<typeof ContainerIdMessageScheme>;
</script>
