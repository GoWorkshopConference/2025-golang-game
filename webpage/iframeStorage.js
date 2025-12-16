// iframeとのlocalStorage連携機能

// iframeのオリジンを取得（セキュリティのため）
function getIframeOrigin(iframe) {
  try {
    return iframe.src ? new URL(iframe.src).origin : window.location.origin;
  } catch (e) {
    return window.location.origin;
  }
}

export function initIframeStorage(iframe) {
  // iframeが読み込まれたらlocalStorageのデータを送信
  iframe.addEventListener("load", () => {
    // localStorageの全データを取得
    const localStorageData = {};
    for (let i = 0; i < localStorage.length; i++) {
      const key = localStorage.key(i);
      if (key) {
        const value = localStorage.getItem(key);
        // 値が存在する場合のみ追加
        if (value !== null) {
          localStorageData[key] = value;
        }
      }
    }
    // iframeのオリジンを取得してセキュアに送信
    const targetOrigin = getIframeOrigin(iframe);
    iframe.contentWindow.postMessage(
      {
        type: "localStorage",
        data: localStorageData,
      },
      targetOrigin,
    );
  });

  // localStorageが変更されたらiframeに通知
  window.addEventListener("storage", (e) => {
    // iframeのオリジンを取得してセキュアに送信
    const targetOrigin = getIframeOrigin(iframe);
    iframe.contentWindow.postMessage(
      {
        type: "localStorage",
        data: { [e.key]: e.newValue },
      },
      targetOrigin,
    );
  });
}

