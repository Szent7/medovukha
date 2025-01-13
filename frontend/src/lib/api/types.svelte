<script lang="ts" context="module">
	import { z } from 'zod';

	export const ContainerIdMessageScheme = z.object({
		Id: z.string()
	});
	export type ContainerIdMessage = z.infer<typeof ContainerIdMessageScheme>;

	export const ContainerBaseInfoScheme = z.array(
		z.object({
			Id: z.string(),
			Names: z.array(z.string()),
			Image: z.string(),
			Ports: z.array(
				z
					.object({
						IP: z.string(),
						PrivatePort: z.number(),
						PublicPort: z.number(),
						Type: z.string()
					})
					.optional()
			),
			Created: z.number(),
			State: z.string(),
			IsMedovukha: z.boolean()
		})
	);
	export type ContainerBaseInfo = z.infer<typeof ContainerBaseInfoScheme>;

	export const ImageBaseInfoScheme = z.array(
		z.object({
			Id: z.string(),
			Tags: z.array(z.string()),
			Size: z.number(),
			Created: z.number()
		})
	);
	export type ImageBaseInfo = z.infer<typeof ImageBaseInfoScheme>;

	export const NetworkBaseInfoScheme = z.array(
		z.object({
			Name: z.string(),
			Id: z.string(),
			Driver: z.string(),
			EnableIPv6: z.boolean(),
			IPAMDriver: z.string(),
			Subnet: z.array(z.string()),
			Gateway: z.array(z.string()),
			Attachable: z.boolean(),
			DockerNetwork: z.boolean()
		})
	);
	export type NetworkBaseInfo = z.infer<typeof NetworkBaseInfoScheme>;

	export const VolumeBaseInfoScheme = z.array(
		z.object({
			Name: z.string(),
			Driver: z.string(),
			Mountpoint: z.string(),
			Created: z.string()
		})
	);
	export type VolumeBaseInfo = z.infer<typeof VolumeBaseInfoScheme>;
</script>
