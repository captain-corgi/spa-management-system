import React, { useState } from 'react';
import './index.css';

interface Item {
  id: number;
  name: string;
  description: string;
}

const mockItems: Item[] = [
  { id: 1, name: 'Item 1', description: 'Description for item 1' },
  { id: 2, name: 'Item 2', description: 'Description for item 2' },
  { id: 3, name: 'Item 3', description: 'Description for item 3' },
  { id: 4, name: 'Item 4', description: 'Description for item 4' },
  { id: 5, name: 'Item 5', description: 'Description for item 5' },
];

const appStyles = {
  container: {
    padding: '1rem',
    maxWidth: '1200px',
    margin: '0 auto',
  },
  header: {
    borderBottom: '1px solid #eee',
    paddingBottom: '1rem',
    marginBottom: '2rem',
  },
  content: {
    display: 'grid',
    gridTemplateColumns: '1fr 1fr',
    gap: '2rem',
  },
  list: {
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
    padding: '1rem',
  },
  listHeader: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    borderBottom: '1px solid #eee',
    paddingBottom: '0.5rem',
    marginBottom: '1rem',
  },
  listItems: {
    listStyle: 'none',
    padding: 0,
    margin: 0,
  },
  listItem: {
    padding: '0.75rem',
    borderBottom: '1px solid #f5f5f5',
    cursor: 'pointer',
    transition: 'background-color 0.2s',
  },
  listItemActive: {
    backgroundColor: '#f0f7ff',
  },
  detail: {
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
    padding: '1rem',
  },
  detailHeader: {
    borderBottom: '1px solid #eee',
    paddingBottom: '0.5rem',
    marginBottom: '1rem',
  },
  detailContent: {
    padding: '1rem 0',
  },
  field: {
    marginBottom: '1rem',
  },
  label: {
    fontWeight: 'bold',
    marginBottom: '0.25rem',
    display: 'block',
  },
  placeholder: {
    textAlign: 'center' as const,
    padding: '2rem',
    color: '#999',
    backgroundColor: '#f9f9f9',
    borderRadius: '8px',
    border: '1px dashed #ddd',
  },
};

export default function App() {
  const [selectedItemId, setSelectedItemId] = useState<number | null>(null);
  
  const selectedItem = selectedItemId 
    ? mockItems.find(item => item.id === selectedItemId) 
    : null;

  return (
    <div style={appStyles.container}>
      <div style={appStyles.header}>
        <h2>uranch Management</h2>
      </div>
      
      <div style={appStyles.content}>
        {/* List View */}
        <div style={appStyles.list}>
          <div style={appStyles.listHeader}>
            <h3>uranch List</h3>
            <button>Add New</button>
          </div>
          
          <ul style={appStyles.listItems}>
            {mockItems.map(item => (
              <li 
                key={item.id} 
                style={{
                  ...appStyles.listItem,
                  ...(selectedItemId === item.id ? appStyles.listItemActive : {})
                }}
                onClick={() => setSelectedItemId(item.id)}
              >
                {item.name}
              </li>
            ))}
          </ul>
        </div>
        
        {/* Detail View */}
        <div style={appStyles.detail}>
          {selectedItem ? (
            <>
              <div style={appStyles.detailHeader}>
                <h3>{selectedItem.name} Details</h3>
              </div>
              
              <div style={appStyles.detailContent}>
                <div style={appStyles.field}>
                  <span style={appStyles.label}>ID:</span>
                  <span>{selectedItem.id}</span>
                </div>
                
                <div style={appStyles.field}>
                  <span style={appStyles.label}>Name:</span>
                  <span>{selectedItem.name}</span>
                </div>
                
                <div style={appStyles.field}>
                  <span style={appStyles.label}>Description:</span>
                  <span>{selectedItem.description}</span>
                </div>
              </div>
            </>
          ) : (
            <div style={appStyles.placeholder}>
              <p>Select an item to view details</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
