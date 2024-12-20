import "./Modal.css";
import { createPortal } from "react-dom";

function Modal({ children }) {
  return createPortal(
    <div className="modal-background">{children}</div>,
    document.getElementById("modal")
  );
}

export { Modal };
