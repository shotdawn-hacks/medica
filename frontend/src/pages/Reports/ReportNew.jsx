import React, { useState } from 'react';

export default function ReportNew() {
  const [typeStart, setTypeStart] = useState('text');
  const [typeEnd, setTypeEnd] = useState('text');
  return (
    <div className="thing">
      <div className="to">
        <div className="consider date">
          <input
            className="dateStart"
            type={typeStart}
            placeholder="Начало"
            onFocus={() => setTypeStart('date')}
            onBlur={() => setTypeStart('text')}
          />
          <input
            className="dateEnd"
            type={typeEnd}
            placeholder="Конец"
            onFocus={() => setTypeEnd('date')}
            onBlur={() => setTypeEnd('text')}
          />
        </div>
        <div className="consider">
          <input
          />
        </div>
      </div>
    </div>
  );
}
