
<script>
import Post from "../components/Post.vue";
import Modal from "../components/Modal.vue";
import { useRoute, useRouter } from 'vue-router';
import { getCookie } from "../utils/cookieUtils";
import { watch } from 'vue';
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			user_profile: null,
			followingModalIsVisible: false,
			followersModalIsVisible: false,
			updateNameModalIsVisible: false,
			isFollowing: false,
			isBanned: false,
			usernameText: '',

			isPersonalProfile: false,

		};
	},
	setup() {
		const route = useRoute();
		const router = useRouter();
		const userId = route.params.userId;

		const identifier = getCookie('identifier');

		watch(() => route.params.userId, (newUserId) => {
			router.go(0);
		});
		return { userId, identifier };

	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;

			this.isPersonalProfile = this.userId == this.identifier;

			try {
				let response = await this.$axios.get("/user/profile/" + this.userId, { headers: { 'Authorization': this.identifier } });
				this.user_profile = response.data;
				this.isFollowing = this.user_profile.IsFollowing;
				this.isBanned = this.user_profile.IsBanned;
				console.log(this.user_profile)
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		handleFollowersToggle() {
			this.followersModalIsVisible = !this.followersModalIsVisible;
		},
		handleFollowingToggle() {
			this.followingModalIsVisible = !this.followingModalIsVisible;
		},
		handleUpdateNameToggle() {
			this.updateNameModalIsVisible = !this.updateNameModalIsVisible;
		},
		deleteImage(photoId) {
			this.user_profile.Photos = this.user_profile.Photos.filter(photo => photo.ID != photoId);
		},
		updateUsername() {
			this.$axios.put("/user/name", { name: this.usernameText }, { headers: { 'Authorization': this.identifier } })
				.then((response) => {
					this.user_profile.Name = this.usernameText;
					this.handleUpdateNameToggle();
					this.refresh();
				})
				.catch((error) => {
					console.log(error);
					window.alert("Username already taken");
					this.handleUpdateNameToggle();
				})
				
				this.usernameText = '';
		},
		toggleFollow() {
			if (this.isFollowing) {
				console.log("delete")
				this.$axios.delete("/user/follow/" + this.userId, { headers: { 'Authorization': this.identifier } })
					.then((response) => {
						this.isFollowing = false;
						this.user_profile.Followers = response.data;
					})
					.catch((error) => {
						console.log(error);
					})
			}
			else {
				this.$axios.post("/user/follow/" + this.userId, {}, { headers: { 'Authorization': this.identifier } })
					.then((response) => {
						this.isFollowing = true;
						this.user_profile.Followers = response.data;
					})
					.catch((error) => {
						console.log(error);
					})
			}
		},
		toggleBan() {
			if (this.isBanned) {
				this.$axios.delete("/user/ban/" + this.userId, { headers: { 'Authorization': this.identifier } })
					.then((response) => {
						this.isBanned = false;
					})
					.catch((error) => {
						console.log(error);
					})
			}
			else {
				window.alert("Are you sure you want to ban this user?")
				this.$axios.post("/user/ban/" + this.userId, {}, { headers: { 'Authorization': this.identifier } })
					.then((response) => {
						this.isBanned = true;
					})
					.catch((error) => {
						console.log(error);
					})
			}
		}
	},
	mounted() {
		console.log("userId" + this.userId)
		this.refresh();
	},
	components: { Post }
}
</script>

<template>
	<AuthWrapper :redirects="true">

		<div v-if="user_profile != null">
			<!-- User Name -->
			<div class="header">
				<Modal :show="followersModalIsVisible" @close="handleFollowersToggle" :users="user_profile.Followers">
					<template v-slot:header>
						<h3>Followers</h3>
					</template>
				</Modal>
				<Modal :show="followingModalIsVisible" @close="handleFollowingToggle" :users="user_profile.Following">
					<template v-slot:header>
						<h3>Following</h3>
					</template>
				</Modal>
				<Modal :show="updateNameModalIsVisible" @close="handleUpdateNameToggle">
					<template v-slot:header>
						<h3>Update Username</h3>
					</template>
					<template v-slot:body>
						<form class="username-form">
							<input type="text" v-model="usernameText" placeholder="New username" />
							<button type="submit" @click.prevent="updateUsername">Update</button>
						</form>
					</template>
				</Modal>
				<div class="top">
					<h1>{{ user_profile.Name }}
						<button v-if="isPersonalProfile" class="" @click="handleUpdateNameToggle">
							<svg class="feather edit">
								<use href="/feather-sprite-v4.29.0.svg#edit" />
							</svg>
						</button>
					</h1>

					<div v-if="!isPersonalProfile">
						<button class="follow-btn" @click="toggleFollow">{{ isFollowing ? "Unfollow" : "Follow" }}</button>
						<button class="ban-btn" @click="toggleBan">{{ isBanned ? "Unban" : "Ban" }}</button>
					</div>
				</div>
				<div class="bottom">
					<h4>{{ user_profile.Photos.length + ((user_profile.Photos.length > 1) ? " posts" : " post") }}</h4>
					<h4 @click="handleFollowingToggle">{{ user_profile.Following.length + " following" }}</h4>
					<h4 @click="handleFollowersToggle">{{ user_profile.Followers.length + ((user_profile.Followers.length >
						1) ? " followers" : " follower"
					) }}</h4>
				</div>
			</div>



			<div class="feed" v-for="post in user_profile.Photos" :key="post.ID">
				<Post :post="post" :identifier="this.identifier" @delete="() => deleteImage(post.ID)" />
			</div>

		</div>

	</AuthWrapper>
</template>
  

  
 
  

<!-- <template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style> -->
