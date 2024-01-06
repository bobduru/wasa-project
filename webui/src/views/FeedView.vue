
<script>
import Post from "../components/Post.vue";
import { getCookie } from "../utils/cookieUtils";

export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			feed: null,
			identifier : getCookie('identifier')
		};
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
                
                // this.$axios.defaults.headers.common['Authorization'] = identifier;
				let response = await this.$axios.get("/stream", { headers: { 'Authorization': this.identifier } })
				this.feed = response.data;
				console.log(response.data);
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh();
	},
	components: { Post }
}
</script>

<template>
	<AuthWrapper :redirects="true">

		<div v-if="feed != null" style="margin-top: 30px;">

			<div class="feed" v-for="post in feed" :key="post.ID">
				<Post :post="post" :identifier="this.identifier"/>
			</div>

		</div>
		<div v-else>
			<p>Looks a bit empty here, you should follow some users</p>
		</div>
	</AuthWrapper>
</template>
  