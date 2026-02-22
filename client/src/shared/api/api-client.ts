interface FetchOptions extends RequestInit {}
const BASE_URL: string = process.env.BASE_URL || "http://localhost:4000";

export async function $api<T>(endpoint: string, options: FetchOptions = {}) {
  const url = `${BASE_URL}${endpoint}`;
  const headers = new Headers(options.headers);
  if (!headers.has("Content-Type")) {
    headers.set("Content-Type", "application/json");
  }
  const config: RequestInit = {
    ...options,
    headers,
  };
  const response = await fetch(url, config);

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(
      errorData.message || `API request failed with status ${response.status}`,
    );
  }

  if (response.status === 204) {
    return {} as T;
  }

  return response.json();
}
