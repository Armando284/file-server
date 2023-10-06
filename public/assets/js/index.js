"strict mode";

const $ = (tag) => {
  if (!tag && typeof tag !== 'string' && tag === '') return null;
  const elements = document.querySelectorAll(tag);
  return elements.length > 1 ? elements : elements[0];
}

const form = $('#multiple-files'),
  cancelBtn = $('#cancel-btn');

cancelBtn.onclick = (e) => {
  e.preventDefault();
  e.stopPropagation();
  form.reset();
}