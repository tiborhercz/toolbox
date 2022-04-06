### Wasm "syscall/js.finalizeRef" not implemented

In tinygo "syscall/js.finalizeRef" is not implemented. To fix the console errors manually implement "syscall/js.finalizeRef"

See https://github.com/tinygo-org/tinygo/issues/1140

```
					// func finalizeRef(v ref)
					"syscall/js.finalizeRef": (v_addr) => {
						// Note: TinyGo does not support finalizers so this is only called
						// for one specific case, by js.go:jsString.
						const id = mem().getUint32(v_addr, true);
						this._goRefCounts[id]--;
						if (this._goRefCounts[id] === 0) {
							const v = this._values[id];
							this._values[id] = null;
							this._ids.delete(v);
							this._idPool.push(id);
						}
					},
```
