#!/bin/bash

# Define ScyllaDB connection parameters
SCYLLA_HOST="127.0.0.1"
SCYLLA_PORT="9042"

# Define the keyspace and table creation CQL commands
CREATE_KEYSPACE_CQL="CREATE KEYSPACE IF NOT EXISTS chatchatgo WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};"

CREATE_TABLE_conversations_by_user_CQL="CREATE TABLE IF NOT EXISTS conversations_by_user (
    sender_id UUID,
    conversation_id TIMEUUID,
    participant_id UUID,
    last_message_id UUID,
    last_message_content TEXT,
    last_message_timestamp TIMESTAMP,
    created_at TIMESTAMP,
    PRIMARY KEY((sender_id), conversation_id)
); WITH CLUSTERING ORDER BY (conversation_id, DESC)"

CREATE_TABLE_messages_CQL="CREATE TABLE IF NOT EXISTS messages (
    message_id TIMEUUID,
    conversation_id UUID,
    content TEXT
    sender_id UUID,
    receiver_id UUID,
    created_at TIMESTAMP,
    PRIMARY KEY(message_id)
);"

CREATE_TABLE_messages_by_conversation_CQL="CREATE TABLE IF NOT EXISTS messages_by_conversation (
    conversation_id UUID,
    message_id TIMEUUID,
    content TEXT
    sender_id UUID,
    receiver_id UUID,
    created_at TIMESTAMP,
    PRIMARY KEY((conversation_id), message_id)
); WITH CLUSTERING ORDER BY (message_id, DESC)"

# Function to execute CQL commands
execute_cql() {
    local cql_command=$1
    echo "exec " $cql_command 
    echo $cql_command | cqlsh $SCYLLA_HOST $SCYLLA_PORT
}

# Create keyspace
echo "Creating keyspace..."
execute_cql "$CREATE_KEYSPACE_CQL"

# Create tables
echo "Creating table conversations_by_user..."
execute_cql "$CREATE_TABLE_conversations_by_user_CQL"

echo "Creating table messages..."
execute_cql "$CREATE_TABLE_messages_CQL"

echo "Creating table messages_by_conversation..."
execute_cql "$CREATE_TABLE_messages_by_conversation_CQL"

echo "Keyspace and table created successfully."