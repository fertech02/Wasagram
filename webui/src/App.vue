<script setup>
import { RouterLink, RouterView } from 'vue-router';
</script>

<script>
export default {
	data() {
		return {

		};
	},
}

</script>

<template>
	<div>
		<header>
			<nav class="navbar navbar-expand-lg navbar-light bg-light">
				<div class="container-fluid">
					<router-link to="/" class="navbar-brand">Wasaphoto</router-link>
					<button class="navbar-toggler" type="button" data-bs-toggle="collapse"
						data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false"
						aria-label="Toggle navigation">
						<span class="navbar-toggler-icon"></span>
					</button>
					<div class="collapse navbar-collapse" id="navbarNav">
						<ul class="navbar-nav">
							<li class="nav-item">
								<router-link to="/session" class="nav-link">Session</router-link>
							</li>
							<li class="nav-item">
								<router-link to="/upload" class="nav-link">Upload</router-link>
							</li>
							<li class="nav-item">
								<router-link to="/logout" class="nav-link">Logout</router-link>
							</li>
						</ul>
					</div>
				</div>
			</nav>
		</header>
	</div>
</template>
