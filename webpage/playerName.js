// プレイヤー名の入力・保存機能

// HTMLタグやスクリプトを除去・エスケープする関数
function sanitizeInput(input) {
  // HTMLタグを除去
  const div = document.createElement("div");
  div.textContent = input;
  return div.textContent || div.innerText || "";
}

// 入力値の検証
function validatePlayerName(name) {
  if (!name || name.trim().length === 0) {
    return { valid: false, error: "名前を入力してください" };
  }
  if (name.length > 20) {
    return { valid: false, error: "名前は20文字以内で入力してください" };
  }
  // 制御文字や特殊な文字をチェック
  if (/[\x00-\x1F\x7F]/.test(name)) {
    return { valid: false, error: "無効な文字が含まれています" };
  }
  return { valid: true };
}

// iframeのオリジンを取得（セキュリティのため）
function getIframeOrigin(iframe) {
  try {
    return iframe.src ? new URL(iframe.src).origin : window.location.origin;
  } catch (e) {
    return window.location.origin;
  }
}

export function initPlayerName(iframe) {
  const playerNameInput = document.getElementById("playerNameInput");
  const saveNameButton = document.getElementById("saveNameButton");

  // ページ読み込み時に保存されている名前を表示
  const savedName = localStorage.getItem("playerName");
  if (savedName) {
    // 保存されている値もサニタイズして表示
    playerNameInput.value = sanitizeInput(savedName);
  }

  // 保存ボタンのクリックイベント
  saveNameButton.addEventListener("click", () => {
    const rawInput = playerNameInput.value.trim();
    const playerName = sanitizeInput(rawInput);
    
    // 検証
    const validation = validatePlayerName(playerName);
    if (!validation.valid) {
      alert(validation.error);
      return;
    }

    // サニタイズされた値を保存
    localStorage.setItem("playerName", playerName);
    
    // iframeのオリジンを取得してセキュアに送信
    const targetOrigin = getIframeOrigin(iframe);
    iframe.contentWindow.postMessage(
      {
        type: "localStorage",
        data: { playerName: playerName },
      },
      targetOrigin,
    );
    alert("名前を保存しました！");
  });

  // Enterキーでも保存できるようにする
  playerNameInput.addEventListener("keypress", (e) => {
    if (e.key === "Enter") {
      saveNameButton.click();
    }
  });
}

