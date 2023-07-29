import type { Task } from '$lib/types';

export const loadTaskList = async function () {
	const taskList = await fetch('http://localhost:80/task/list', { method: 'GET' })
		.then((response) => {
			if (!response.ok) {
				return new Error();
			}
			return response.json();
		})
		.catch(() => {
			console.error('failed to fetch Task List');
		});

	// TODO: アサーション使わないで済む方法
	return taskList as Task[];
};

export const addTask = async function (title: string) {
	const requestBody = {
		title
	};
	fetch('http://localhost:80/task/add', {
		method: 'POST',
		headers: {
			Accept: 'application/json',
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(requestBody)
	})
		.then((response) => {
			if (!response.ok) return new Error();
		})
		.catch((error) => {
			console.error('failed to Add Task:', error);
		});
};

export const removeTask = async function (id: number | string) {
	const requestBody = {
		id: Number(id)
	};
	fetch('http://localhost:80/task/remove', {
		method: 'POST',
		headers: {
			Accept: 'application/json',
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(requestBody)
	})
		.then((response) => {
			if (!response.ok) return new Error();
		})
		.catch((error) => {
			console.error('failed to remove Task:', error);
		});
};
