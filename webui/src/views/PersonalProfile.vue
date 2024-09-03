<script>
export default {
	methods: {
		async redirect() {
		    this.$router.replace(`/users/${sessionStorage.getItem('token')}/profile`);
		},
	},
	mounted() {
		this.redirect()
	}
}
</script>

<template>
	<span></span>
</template>

<style>
</style>