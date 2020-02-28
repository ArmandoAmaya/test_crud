<template>
	<div>
		<b-container class="mt-5">
			<b-card>
				<b-card-title>LISTADO de usuarios</b-card-title>
				<b-table :fields="fields" striped hover :items="items">
					<template v-slot:cell(action)="data">
						<router-link :to="'/usuarios/editar/'+data.item.id" v-b-tooltip.hover title="Editar">
							<b-icon icon="pencil" variant="primary"></b-icon>
						</router-link> 
						&nbsp;
						<a href="javascript:void(0)" @click.prevent="deleteUser(data.item.id)" v-b-tooltip.hover title="Eliminar">
							<b-icon icon="trash" variant="danger"></b-icon>
						</a> 
						
					</template>
				</b-table>
				<b-button variant="outline-primary" to="/usuarios/crear">Crear nuevo</b-button>
			</b-card>
		</b-container>
	</div>
</template>

<script>
//import axios from "axios";
import axios from "axios";
export default {
	name: "Home",
	data(){
		return {
			fields: [
				{ key: 'id', label: 'ID' },
				{ key: 'name', label: 'Nombre' },
				{ key: 'email', label: 'Email' },
				{ key: 'regitered_date', label: 'Fecha de registro' },
				{ key: 'action', label: 'Acción' },
			],
			items: []
		}
	},
	methods: {
		deleteUser(id_user){
			var __data = new FormData();
			__data.append('id_user', id_user)

			axios
			.post(this.$store.state.api + "usuarios/delete", __data, {
				headers: {
					"Content-Type": "multipart/form-data"
				}
			})
			.then(resp => console.log(resp));

			var parseData = [];

			this.items.forEach(data => {
				if (id_user != data.id)
					parseData.push(data);				
			});

			this.items = parseData;
		}
	},
	mounted(){
		axios
		.get(this.$store.state.api + "usuarios", {
			headers: {
				"Content-Type": "multipart/form-data"
			}
		})
		.then(resp => {
			var data = resp.data;
			var parseData = [];
			data.forEach(d => {
				parseData.push({
					id: d.id_user,
					name: d.name,
					email: d.email,
					regitered_date: d.str_register_date,
					action: ''
				});
			});

			this.items = parseData;
		});
	}
};
</script>

<style lang="scss" scoped></style>
