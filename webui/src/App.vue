<!-- <script setup>
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { getCookie, deleteCookie } from './utils/cookieUtils';
import { onMounted, watch } from 'vue';
import { provide, reactive, ref } from 'vue';
import Modal from "./components/Modal.vue";
// import { provide, ref } from 'vue'

// const router = useRouter();
// const loggedIn = ref(!!getCookie('identifier'))

// function setLoggedIn(value) {
// 	loggedIn.value = value
// }

// provide('loggedIn', {
// 	loggedIn,
// 	setLoggedIn
// })
</script> -->
<script>
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { getCookie, deleteCookie } from './utils/cookieUtils';
import { onMounted, watch } from 'vue';
import { provide, reactive, ref } from 'vue';
import Modal from "./components/Modal.vue";

export default {
	data() {
		return {
			users: [],
			searchModalIsVisible: false,
			identifier: getCookie('identifier'),
		}
	},
	setup() {
		const router = useRouter();
		const loggedIn = ref(!!getCookie('identifier'))

		function setLoggedIn(value) {
			loggedIn.value = value
		}
		

		provide('loggedIn', {
			loggedIn,
			setLoggedIn
		})
		return {
			loggedIn,
			setLoggedIn, 
			router
		}
	},
	methods: {
		logout() {
			
			this.setLoggedIn(false);
			deleteCookie('identifier');
			this.router.push('/login');
		},
		navigateToAccount() {

			this.router.push('/user/'+ getCookie('identifier'));
		},
		async refresh() {
			try {
				let response = await this.$axios.get("/users");
				this.users = response.data;
				// console.log(this.users)
			}
			catch (e) {
				this.errormsg = e.toString();
			}
		},
		handleSearchModalToggle() {
			this.searchModalIsVisible = !this.searchModalIsVisible;
		},
	},
	mounted() {
		this.refresh();
	},
};
</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse"
			data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<!-- <h6
						class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6> -->
					<Modal :show="searchModalIsVisible" @close="handleSearchModalToggle" :users="users">
						<template v-slot:header>
							<h3>Users</h3>
						</template>
					</Modal>
					<ul class="nav flex-column">
						<li class="nav-item" v-if="loggedIn">
							<button class="nav-link" @click="handleSearchModalToggle">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#search" />
								</svg>
								Search
							</button>
						</li>
						<li class="nav-item" v-if="loggedIn">
							<RouterLink to="/" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#list" />
								</svg>
								Feed
							</RouterLink>
						</li>
						
						<li class="nav-item" v-if="loggedIn">
							<a class="nav-link" @click="navigateToAccount">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#user" />
								</svg>
								Account
							</a>
						</li>

						<li class="nav-item" v-if="loggedIn">
							<RouterLink class="nav-link" to="/create">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#plus-circle" />
								</svg>
								Create Post
							</RouterLink>
						</li>
						<li class="nav-item" v-if="loggedIn">
							<a class="nav-link" @click="logout">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#log-out" />
								</svg>
								Logout
							</a>
						</li>
						<li class="nav-item" v-else>
							<RouterLink to="/login" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#log-in" />
								</svg>
								Login
							</RouterLink>
						</li>
					</ul>

				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style></style>
