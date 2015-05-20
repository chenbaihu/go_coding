#include "echo.h"

#include <muduo/base/Logging.h>

#include <boost/bind.hpp>

// using namespace muduo;
// using namespace muduo::net;

EchoServer::EchoServer(muduo::net::EventLoop* loop,
                       const muduo::net::InetAddress& listenAddr)
  : server_(loop, listenAddr, "EchoServer")
{
  server_.setConnectionCallback(
      boost::bind(&EchoServer::onConnection, this, _1));
  server_.setMessageCallback(
      boost::bind(&EchoServer::onMessage, this, _1, _2, _3));
}

void EchoServer::Start(int thread_count)
{
  server_.setThreadNum(thread_count);
  server_.start();
}

void EchoServer::onConnection(const muduo::net::TcpConnectionPtr& conn)
{
  LOG_INFO << "EchoServer - " << conn->peerAddress().toIpPort() << " -> "
           << conn->localAddress().toIpPort() << " is "
           << (conn->connected() ? "UP" : "DOWN");
}

void EchoServer::onMessage(const muduo::net::TcpConnectionPtr& conn,
                           muduo::net::Buffer* buf,
                           muduo::Timestamp time)
{
  const static uint32_t head_len = 4;
  while (buf->readableBytes()>head_len) { 
      const char* data = buf->peek();
      uint32_t len = ntohl(*(uint32_t*)(data));
     
      if (buf->readableBytes() >= (uint32_t)(head_len + len)) {
          muduo::string msg(data, head_len + len);
          LOG_INFO << conn->name() << " echo " << msg.size() << " bytes, " << "data received at " << time.toString(); 
          buf->retrieve(head_len + len);
          conn->send(msg);
      } else {
          LOG_INFO << conn->name() << " echo len=" << len << "\treadAbleBytes=" << buf->readableBytes(); 
          break;
      }
  }
  //muduo::string msg(buf->retrieveAllAsString());
  //LOG_INFO << conn->name() << " echo " << msg.size() << " bytes, "
  //         << "data received at " << time.toString();
  //conn->send(msg);
}

