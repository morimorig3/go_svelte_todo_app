import { addTask, removeTask } from '$lib/server/task';
import type { Actions } from './$types';

export const actions: Actions = {
	add: async ({ request }) => {
		const data = await request.formData();
		const taskTitle = data.get('taskTitle') as string;
		await addTask(taskTitle);
	},
	remove: async ({ request }) => {
		const data = await request.formData();
		const taskId = data.get('taskId') as string;
		await removeTask(taskId);
	}
};
