let socket;

// Function to generate a random Materialize color class
function getRandomColorClass() {
  const colors = [
    "red", "pink", "purple", "deep-purple", 
    "indigo", "blue", "light-blue", "cyan", 
    "teal", "green", "light-green", "lime", 
    "yellow", "amber", "orange", "deep-orange", 
    "brown", "grey", "blue-grey"
  ];
  const randomIndex = Math.floor(Math.random() * colors.length);
  return colors[randomIndex];
}

// Object to store sender-to-color mapping
const senderColors = {};

$(document).ready(function () {
  // Connect button
  $("#connect-btn").click(function () {
    if (socket) {
      M.toast({ html: "Already connected!" });
      return;
    }
    
    // Open WebSocket connection
    socket = new WebSocket("/chat");

    socket.onopen = function () {
      M.toast({ html: "Connected to the chat!" });
    };

    socket.onmessage = function (event) {
      const message = event.data;

      if (message.includes(":")) {
        // Split the message at the first colon
        const [sender, ...messageParts] = message.split(":");
        const actualMessage = messageParts.join(":").trim();

        // Assign a color to the sender if not already assigned
        if (!senderColors[sender]) {
          senderColors[sender] = getRandomColorClass();
        }

        const colorClass = senderColors[sender];

        // Append message with sender info and profile symbol
        $("#chat-box").append(`
        <li class="collection-item avatar custom-message-item">
            <i class="material-icons circle ${colorClass}"></i>
            <div class="message-content">
                <span class="title sender-name"><strong>${sender}</strong></span>
                <p class="message-text">${actualMessage}</p>
            </div>
        </li>
        `);
      } else {
        // Handle messages without colons as highlighted system messages
        $("#chat-box").append(`
        <li class="collection-item highlighted-message-item">
            <div class="highlighted-message-content">
                <span class="highlighted-message">${message}</span>
            </div>
        </li>
        `);
      }

      // Scroll to the bottom of the chat box
      const chatBox = document.getElementById("chat-box");
      chatBox.scrollTop = chatBox.scrollHeight;
    };

    socket.onerror = function () {
      M.toast({ html: "Error connecting to the server!" });
    };

    socket.onclose = function () {
      M.toast({ html: "Connection closed." });
      socket = null;
    };
  });

  // Send message on Enter
  $("#message-input").keypress(function (e) {
    if (e.which == 13 && socket && socket.readyState === WebSocket.OPEN) {
      const message = $(this).val();
      if (message.trim() !== "") {
        socket.send(message);
        $(this).val(""); // Clear input
      }
    }
  });
});
