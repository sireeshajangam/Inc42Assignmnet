// Sample unit test using Jest
import { fetchData } from '../index';

test('fetchData function works', async () => {
  const data = await fetchData();
  expect(data).toBeDefined();
  expect(Array.isArray(data)).toBe(true);
});
