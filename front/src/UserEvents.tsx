import { useState, type FormEvent } from "react";

export function UserEvents() {
  const [events, setEvents] = useState([])

  const parseDate = (d: any): string => {
    return new Date(d).toISOString()
  }

  const testEndpoint = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const form = e.currentTarget;
      const formData = new FormData(form);

      const p = new URLSearchParams()
      p.append("user_id", String(formData.get("user_id")))
      p.append("from", parseDate(formData.get("from")))
      p.append("till", parseDate(formData.get("till")))

      const apiBaseURL = 'http://localhost:8080';

      const url = `${apiBaseURL}/api/events?${p.toString()}`
      const res = await fetch(url, { method: "GET" });

      const data = await res.json();
      setEvents(data.events)
    } catch (error) { }
  };

  return (
    <div className="user-events">
      <h1>Find events</h1>
      <form onSubmit={testEndpoint} className="endpoint-row">
        <div className="input-container">
          <div>
            <span>User id</span>
            <input type="text" name="user_id" defaultValue="" className="id-input" placeholder="Input user id" />
          </div>
          <div>
            <span>From</span>
            <input type="date" name="from" value="2025-10-25" className="id-input" />
          </div>
          <div>
            <span>Till</span>
            <input type="date" name="till" value="2025-10-26" className="id-input" />
          </div>
        </div>

        <button type="submit" className="search-button">
          Search
        </button>
      </form>

      <div className="response">
        <table className="results-table">
          <thead>
            <tr>
              <th>User ID</th>
              <th>Action</th>
              <th>Occured At</th>
              <th>Metadata</th>
            </tr>
          </thead>
          <tbody>
            {events?.map((e: any) => <tr>
              <td>
                {e.user_id}
              </td>
              <td>
                {e.action}
              </td>
              <td>
                {e.occured_at}
              </td>
              <td>
                {JSON.stringify(e.metadata)}
              </td>
            </tr>)}
          </tbody>
        </table>
      </div>
    </div>
  );
}
