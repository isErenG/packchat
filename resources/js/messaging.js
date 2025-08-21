document.addEventListener("DOMContentLoaded", () => {
  const session_id = Math.random().toString(36).substring(2, 15);
  console.log(session_id);
  const WS_URL = location.origin.replace(/^http/, "ws") + "/ws";
  const ws = new WebSocket(WS_URL);
  ws.binaryType = "arraybuffer";

  const messages = document.getElementById("messages");
  const form = document.querySelector(".chat form");
  const input = form.querySelector("input");

  ws.onmessage = (ev) => {
    try {
      const m = window.msgpack.deserialize(new Uint8Array(ev.data)); // {A,B,C}
      const text = esc(m.B || "");

      if (session_id == m.A) {
        messages.insertAdjacentHTML(
          "beforeend",
          `<div class="msg-right">${text}</div>`,
        );
      } else {
        messages.insertAdjacentHTML(
          "beforeend",
          `<div class="msg">${text}</div>`,
        );
      }

      messages.scrollTop = messages.scrollHeight;
    } catch (e) {
      console.error("Bad MsgPack:", e);
    }
  };

  form.addEventListener("submit", (e) => {
    e.preventDefault();
    const text = input.value.trim();
    if (!text || ws.readyState !== WebSocket.OPEN) return;
    ws.send(
      window.msgpack.serialize({
        A: session_id,
        B: text,
        C: Date.now(),
      }),
    );
    input.value = "";
  });

  function esc(s) {
    return String(s).replace(
      /[&<>"']/g,
      (c) =>
        ({
          "&": "&amp;",
          "<": "&lt;",
          ">": "&gt;",
          '"': "&quot;",
          "'": "&#39;",
        })[c],
    );
  }
});
