<script lang="ts" context="module">
	import axios from 'axios';
	import {
		ContainerBaseInfoScheme,
		ImageBaseInfoScheme,
		NetworkBaseInfoScheme,
		VolumeBaseInfoScheme,
		type ContainerBaseInfo,
		type ContainerIdMessage,
		type ImageBaseInfo,
		type NetworkBaseInfo,
		type VolumeBaseInfo
	} from './types.svelte';

	export async function PingBack() {
		axios.get('http://localhost:10015/ping').then(function (response) {
			console.log(response.data);
		});
	}

	export async function GetImageList() {
		try {
			const response = await axios.get('http://localhost:10015/rest/v1/getimagelist');
			const images: ImageBaseInfo = ImageBaseInfoScheme.parse(response.data);
			//console.log("responseData:" + containers);
			return images;
		} catch (error) {
			console.error('Error GET ContainerList:', error);
		}
	}

	export async function GetNetworkList() {
		try {
			const response = await axios.get('http://localhost:10015/rest/v1/getnetworklist');
			const networks: NetworkBaseInfo = NetworkBaseInfoScheme.parse(response.data);
			//console.log("responseData:" + containers);
			return networks;
		} catch (error) {
			console.error('Error GET ContainerList:', error);
		}
	}

	export async function GetVolumeList() {
		try {
			const response = await axios.get('http://localhost:10015/rest/v1/getvolumelist');
			const volumes: VolumeBaseInfo = VolumeBaseInfoScheme.parse(response.data);
			//console.log("responseData:" + containers);
			return volumes;
		} catch (error) {
			console.error('Error GET ContainerList:', error);
		}
	}

	export async function GetContainerList() {
		try {
			const response = await axios.get('http://localhost:10015/rest/v1/getcontainerlist');
			const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(response.data);
			//console.log("responseData:" + containers);
			return containers;
		} catch (error) {
			console.error('Error GET ContainerList:', error);
		}
	}

	export async function StartContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/startcontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST StartContainerById:', error);
		}
	}

	export async function PauseContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/pausecontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST StartContainerById:', error);
		}
	}

	export async function UnpauseContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/unpausecontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST StartContainerById:', error);
		}
	}

	export async function RestartContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/restartcontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST RestartContainerById:', error);
		}
	}

	export async function StopContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/stopcontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST StopContainerById:', error);
		}
	}

	export async function KillContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/killcontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST StartContainerById:', error);
		}
	}

	export async function RemoveContainerById(id: string) {
		const container: ContainerIdMessage = {
			Id: id
		};
		try {
			const response = await axios.post(
				'http://localhost:10015/rest/v1/removecontainerbyid',
				container
			);
			/*const containers: ContainerBaseInfo = ContainerBaseInfoScheme.parse(
                response.data,
            );
            console.log("responseData:" + containers);
            return containers;*/
		} catch (error) {
			console.error('Error POST StartContainerById:', error);
		}
	}
</script>
