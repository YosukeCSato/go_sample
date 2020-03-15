<template>
	<div id="app">
		<h1 class="greeting">あいさつ:{{ msg }}</h1>
		<Content />
	</div>
</template>

<script lang="ts">

export default {
	data() {
		return {
			msg: "こんにちは"
		}
	}
}
</script>

<style>
.greeting {
	color: red;
}
</style>
