// iframeとのlocalStorage連携機能

export function initIframeStorage(iframe) {
  // iframeが読み込まれたらlocalStorageのデータを送信
  iframe.addEventListener("load", () => {
    // localStorageの全データを取得
    const localStorageData = {};
    for (let i = 0; i < localStorage.length; i++) {
      const key = localStorage.key(i);
      localStorageData[key] = localStorage.getItem(key);
    }
    // iframeにpostMessageで送信
    iframe.contentWindow.postMessage(
      {
        type: "localStorage",
        data: localStorageData,
      },
      "*",
    );
  });

  // localStorageが変更されたらiframeに通知
  window.addEventListener("storage", (e) => {
    iframe.contentWindow.postMessage(
      {
        type: "localStorage",
        data: { [e.key]: e.newValue },
      },
      "*",
    );
  });
}

