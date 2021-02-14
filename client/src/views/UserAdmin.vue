<template>
  <div class="container">
    <h1>User Admin</h1>
    <CreateUser />
    <h2>User list</h2>
    <DeleteUser v-for="user in users" v-bind:key="user.name" :name="user.name"/>
    <h4 v-if="users.length === 0">There is no users!</h4>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import CreateUser from '@/components/CreateUser.vue';
import DeleteUser from '@/components/DeleteUser.vue';
import { userService } from '@/services';
import { UserResponse } from '@/services/models';

export default defineComponent({
  name: 'UserAdmin',
  components: {
    CreateUser,
    DeleteUser,
  },
  data(): { users: UserResponse[] } {
    return {
      users: [],
    };
  },
  mounted() {
    userService.getAll()
      .then((users) => {
        this.users = users;
      });
  },
});
</script>

<style scoped>
.container {
  bottom: 0px;
  width: 100vw;
  padding: 20px;
  margin-top: 40px;

  display: flex;
  flex-direction: column;
}

.container > h1 {
  width: 100%;
}
</style>
