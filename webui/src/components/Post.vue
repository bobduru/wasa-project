<script>

function userHasLiked(post, identifier) {
    return post.Likes.some(like => like.UserID == identifier);
}

export default {
    props: ["post", "identifier"],
    data: function () {
        return {
            liked: userHasLiked(this.post, this.identifier),
            newComment: ''
        };
    },
    methods: {
        toggleLike() {
            // console.log(this.identifier)
            // this.liked = !this.liked;
            // console.log("Like toggled");
            //if liked send delete request to photo/like/identifier else post, if successfull update liked
            if (this.liked) {
                this.$axios.delete("/photo/like/" + this.post.ID, { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.post.Likes = response.data;
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
            else {
                this.$axios.post("/photo/like/" + this.post.ID, {}, { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.post.Likes = response.data;
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
        },
        submitComment() {
            const payload = {
                comment: this.newComment
            };
            this.$axios.post('/photo/comment?photoId=' + this.post.ID, payload, {
                headers: { 'Authorization': this.identifier }
            })
                .then((response) => {
                    console.log(response);
                    this.post.Comments.push(response.data);
                    this.newComment = ''; // Reset comment input
                })
                .catch((error) => {
                    console.log(error);
                });
        },
        deleteComment(id) {
            console.log(this.post)
            this.$axios.delete('/photo/comment?commentId=' + id, {
                headers: { 'Authorization': this.identifier }
            })
                .then((response) => {
                    console.log(response);
                    this.post.Comments = this.post.Comments.filter(comment => comment.ID != id);
                })
                .catch((error) => {
                    console.log(error);
                });
        }
    }
}
</script>

<template>
    <div v-if="post != null" class="post-container">
        <div class="username-container">
            <RouterLink :to="{ path: 'user/' + post.UserID }" replace>
                <p>{{ post.UserName }}</p>
                <!-- <p>{{ this.test}}</p> -->
            </RouterLink>
        </div>
        <div class="image-container">
            <img :src="'http://localhost:3000/images/' + post.FileName" :alt="post.FileName">
        </div>
        <div class="likes-container">
            <p>
                <button class="like-btn" @click="toggleLike">
                    <svg class="feather" :class="this.liked ? 'liked' : ''">
                        <use href="/feather-sprite-v4.29.0.svg#heart" />
                    </svg>
                </button>
                {{ post.Likes.length + ((post.Likes.length > 1) ? " likes" : " like") }}
            </p>
        </div>
        <div class="comments-container">
            <ul>
                <li class="comment" v-for="comment in post.Comments" :key="comment.ID">
                    <div class="comment-top">

                        <div>
                            <strong>{{ comment.UserName }}</strong> - <small>{{ new
                                Date(comment.CreateTime).toLocaleString()
                            }}</small>
                        </div>
                        <button class="like-btn" @click="() => deleteComment(comment.ID)" v-if="comment.UserID == this.identifier">
                            <svg class="feather" :class="this.liked ? 'liked' : ''">
                                <use href="/feather-sprite-v4.29.0.svg#trash" />
                            </svg>
                        </button>
                    </div>
                    {{ comment.Text }}

                </li>
            </ul>
            <form class="comment-input-container" @submit.prevent="submitComment">
                <input type="text" v-model="newComment" placeholder="Comment" />
                <button type="submit">Post</button>
            </form>
        </div>
    </div>
</template>

<style>

</style>